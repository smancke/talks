import java.util.concurrent.CountDownLatch;
    
public class ThreadExample {

    public static void main(String[] args) throws Exception {

        int N = 1000000;
        
        CountDownLatch doneSignal = new CountDownLatch(N);

        long start = System.currentTimeMillis();
        for (int i=0; i<N; i++) {
            new Thread() {
                public void run() {
                    doneSignal.countDown();
                }            
            }.start();            
        }
        doneSignal.await();
        System.out.println("done in "+ (System.currentTimeMillis()-start) + "ms");
    }
}
