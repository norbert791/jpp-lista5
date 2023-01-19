package PhilosopherProblemJava;

import java.text.MessageFormat;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class Main {

  public static void main(String[] args) {
    final int numberOfPhilosophers = 3;
    final long sleepTimeMs = 500;
    final int timesEaten = 5;
    final List<Fork> forks = new ArrayList<>();
    final List<Philosopher> philosophers = new ArrayList<>();
    var threadPool = Executors.newFixedThreadPool(numberOfPhilosophers);

    for (int i = 0; i < numberOfPhilosophers; i++) {
      forks.add(new Fork());
    }

    for (int i = 0; i < numberOfPhilosophers; i++) {
      philosophers.add(new Philosopher(forks.get(i % numberOfPhilosophers), forks.get((i + 1) % numberOfPhilosophers), sleepTimeMs));
    }
    
    System.out.println("Philosophers are eating their meals.");

    for (var phil : philosophers) {
      threadPool.execute(phil);
    }

    boolean work = true;

    while (work) {
      try {
        Thread.sleep(sleepTimeMs);
        work = false;
        for (var phil : philosophers) {
          if (phil.getTimesEaten() < timesEaten) {work = true;}
        }
        
      } catch(InterruptedException e) {
        System.out.println("Main thread interrupted");
      }
    }

    System.out.println("All the philosophers have eaten enough times. Closing...");
    for (var phil : philosophers) {
      phil.stopEating();
    }
    threadPool.shutdown();
    try {
      if(!threadPool.awaitTermination(sleepTimeMs * (numberOfPhilosophers + 1), TimeUnit.MILLISECONDS)) {
        System.out.println("One of the thread is considered deadlocked. Forcing shutdown...");
        threadPool.shutdownNow();
      }
    } catch(InterruptedException e) {
      System.out.println("Interrupted during leaving");
    }
    
    System.out.println("Philosophers have finished eating");
    System.out.println("Summary:");
    for (var phil : philosophers) {
      System.out.println(MessageFormat.format("Philosopher no {0} has finished with {1} meals eaten", phil.getId(), phil.getTimesEaten()));
    }
  }
}
