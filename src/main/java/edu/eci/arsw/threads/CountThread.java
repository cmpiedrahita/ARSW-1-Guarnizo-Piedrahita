/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.threads;

/**
 *
 * @author hcadavid
 */
public class CountThread extends Thread {

    private int A;
    private int B;

    public CountThread(int A, int B){
        if(A < B){
        this.A = A;
        this.B = B;
        }
        else{
            System.out.println("El intervalo es incorrecto");
        }

    }

    @Override
    public void run() {
        printNumbers();
    }

    public void printNumbers(){
        for(int i = this.A; i<this.B; i++){
            System.out.println(i);
        }
    }
    
}
