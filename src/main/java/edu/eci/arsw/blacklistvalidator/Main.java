package edu.eci.arsw.blacklistvalidator;

import java.util.List;

public class Main {

    public static void main(String a[]){
        System.out.println("EVALUACION DE DESEMPEÑO");

        int nucleos = Runtime.getRuntime().availableProcessors();
        System.out.println("Nucleos del sistema: " + nucleos);
        System.out.println();

        HostBlackListsValidator hblv = new HostBlackListsValidator();
        String ipDispersa = "202.24.34.55";

        int[] configs = {1, nucleos, nucleos*2, 50, 100};
        String[] nombres = {
                "1 hilo",
                nucleos + " hilos (nucleos)",
                (nucleos*2) + " hilos (2x nucleos)",
                "50 hilos",
                "100 hilos"
        };

        System.out.println("IP de prueba: " + ipDispersa);

        for (int i = 0; i < configs.length; i++) {
            System.out.println();
            System.out.println("PRUEBA " + (i+1) + ": " + nombres[i]);
            printLine(30);

            long inicio = System.currentTimeMillis();
            List<Integer> resultados = hblv.checkHost(ipDispersa, configs[i]);
            long fin = System.currentTimeMillis();

            long tiempo = fin - inicio;

            System.out.println("Tiempo: " + tiempo + " ms");
            System.out.println("Encontrado en: " + resultados.size() + " listas");
            System.out.println("Listas: " + resultados);

            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println();
        System.out.println("Todas las pruebas completadas!");
        System.out.println();
        System.out.println("Abre jVisualVM para ver:");
        System.out.println("   - Uso de CPU en cada prueba");
        System.out.println("   - Consumo de memoria");
        System.out.println("   - Threads activos");
    }

    private static void printLine(int length) {
        for (int i = 0; i < length; i++) {
            System.out.print("-");
        }
        System.out.println();
    }
}