/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.blacklistvalidator;

import edu.eci.arsw.spamkeywordsdatasource.HostBlacklistsDataSourceFacade;
import java.util.LinkedList;
import java.util.List;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.logging.Level;
import java.util.logging.Logger;

public class HostBlackListsValidator {

    private static final int BLACK_LIST_ALARM_COUNT = 5;


    public List<Integer> checkHost(String ipaddress, int n) {
        if (n <= 0) {
            throw new IllegalArgumentException("El número de hilos debe ser positivo: " + n);
        }


        AtomicInteger globalOcurrencesCount = new AtomicInteger(0);
        AtomicBoolean stopSignal = new AtomicBoolean(false);

        HostBlacklistsDataSourceFacade skds = HostBlacklistsDataSourceFacade.getInstance();
        int totalServers = skds.getRegisteredServersCount();

        if (n > totalServers) {
            n = totalServers;
        }

        BlackListSearchThread[] threads = new BlackListSearchThread[n];

        int segmentSize = totalServers / n;
        int remainder = totalServers % n;

        int start = 0;
        int totalCheckedLists = 0;
        List<Integer> allFoundServers = new LinkedList<>();

        for (int i = 0; i < n; i++) {
            int end = start + segmentSize - 1;
            if (i < remainder) {
                end++;
            }

            if (end >= totalServers) {
                end = totalServers - 1;
            }

            threads[i] = new BlackListSearchThread(
                    ipaddress,
                    start,
                    end,
                    skds,
                    globalOcurrencesCount,
                    stopSignal
            );
            threads[i].start();
            start = end + 1;
        }
        for (int i = 0; i < n; i++) {
            try {
                threads[i].join();

                totalCheckedLists += threads[i].getCheckedServersCount();
                allFoundServers.addAll(threads[i].getFoundServers());

            } catch (InterruptedException e) {
                LOG.log(Level.SEVERE, "Hilo interrumpido durante join", e);
                Thread.currentThread().interrupt();
            }
        }

        if (globalOcurrencesCount.get() >= BLACK_LIST_ALARM_COUNT) {
            skds.reportAsNotTrustworthy(ipaddress);
        } else {
            skds.reportAsTrustworthy(ipaddress);
        }

        LOG.log(Level.INFO, "Checked Black Lists:{0} of {1}",
                new Object[]{totalCheckedLists, totalServers});

        return allFoundServers;
    }

    private static final Logger LOG = Logger.getLogger(HostBlackListsValidator.class.getName());
}
