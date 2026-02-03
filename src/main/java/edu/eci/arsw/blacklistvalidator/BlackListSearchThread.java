package edu.eci.arsw.blacklistvalidator;

import java.util.List;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;
import edu.eci.arsw.spamkeywordsdatasource.HostBlacklistsDataSourceFacade;

public class BlackListSearchThread extends Thread {

    private String ipAddress;
    private int startServer;
    private int endServer;
    private HostBlacklistsDataSourceFacade facade;
    private AtomicInteger globalCounter;
    private AtomicBoolean stopSignal; // Cambiado a AtomicBoolean
    private List<Integer> foundServers;
    private int checkedServersCount;

    public BlackListSearchThread(String ipAddress, int startServer, int endServer,
                                 HostBlacklistsDataSourceFacade facade,
                                 AtomicInteger globalCounter,
                                 AtomicBoolean stopSignal) { // Cambiado a AtomicBoolean
        this.ipAddress = ipAddress;
        this.startServer = startServer;
        this.endServer = endServer;
        this.facade = facade;
        this.globalCounter = globalCounter;
        this.stopSignal = stopSignal;
        this.foundServers = new java.util.ArrayList<>();
        this.checkedServersCount = 0;
    }

    @Override
    public void run() {
        for (int server = startServer; server <= endServer && !stopSignal.get(); server++) {
            checkedServersCount++;

            if (facade.isInBlackListServer(server, ipAddress)) {
                foundServers.add(server);
                globalCounter.incrementAndGet();

                if (globalCounter.get() >= 5) {
                    stopSignal.set(true);
                    break;
                }
            }
        }
    }

    public List<Integer> getFoundServers() {
        return foundServers;
    }

    public int getOccurrencesCount() {
        return foundServers.size();
    }

    public int getCheckedServersCount() {
        return checkedServersCount;
    }
}