package PhilosopherProblemJava;

import java.text.MessageFormat;

public class Philosopher implements Runnable{ 

  public Philosopher(Fork fork1, Fork fork2, long sleepTimeMs) {
    this.smallerFork = fork1;
    this.largerFork = fork2;
    if (smallerFork.getId() > largerFork.getId()) {
      var temp = smallerFork;
      smallerFork = largerFork;
      largerFork = temp;
    }
    id = generateId();
    this.sleepTimeMs = sleepTimeMs;
    System.out.println(MessageFormat.format(
      "Philosopher {0} created with forks {1} and {2}",
      this.id,
      fork1.getId(),
      fork2.getId()
    ));
  }

  public int getTimesEaten() {return timesEaten;}
  public int getId() {return id;}
  public void stopEating() {
    synchronized(this.working) {
      this.working = false;
    }
  }

  @Override
  public void run() {
    try {
      boolean privWorking = true;
      synchronized(working) {
        working = true;
      }

      while (privWorking) {
        synchronized(smallerFork) {
          System.out.println(MessageFormat.format(
            "Philosopher no. {0} picks up fork {1}", id, smallerFork.getId()));
          synchronized(largerFork) {
            System.out.println(MessageFormat.format(
              "Philosopher no. {0} picks up fork {1}", id, largerFork.getId()));
              Thread.sleep(sleepTimeMs);
            System.out.println(MessageFormat.format(
              "Philosopher no. {0} has finished eating for the {1} time. Forks are being put down.", id, this.timesEaten));
              timesEaten++;
            //largerFork.notify();  
          }
          //smallerFork.notify();
          //smallerFork.wait();
        }
        synchronized(working) {
          privWorking = working;
        }
        /*
         * Note: We assume that philosopher take relevant breaks between meals.
         * If Philosophers are allowed to use forks immidiately after putting them down
         * then additional synchronization may be needed to prevent starvation.
         * 
         * Uncomment notify/wait calls from the code for alternative approach.
         */
        System.out.println(MessageFormat.format("Philosopher {0} is thinking", id));
        Thread.sleep(Math.max(10, sleepTimeMs));
        System.out.println(MessageFormat.format("Philosopher {0} wants to eat", id));
        
        /*synchronized(smallerFork) {
          smallerFork.notify();
        }
        synchronized(largerFork) {
          largerFork.notify();
        }*/
      }
    } catch(InterruptedException e) {
      System.out.print("Thread interrupted, exiting...");
    }
  }
 
  private Fork smallerFork;
  private Fork largerFork;
  private final int id;
  private final long sleepTimeMs;
  private static int idPool = 0;
  private volatile int timesEaten = 0;
  private Boolean working = false;
  private static synchronized int generateId() {return idPool++;}
  
}