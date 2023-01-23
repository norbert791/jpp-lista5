with Ada.Text_IO;
use Ada.Text_IO;

package body Philosopher is

  protected body MyMutex is
      entry lock() is
        when not locked is
      begin
        locked := True;
      end lock;
      
      entry unlock() is
        when locked is
      begin
        locked := False;
      end unlock;
    end MyMutex;
  
  task body RunPhilo is
    philo : PhilosopherPtr;
    work : Boolean;

  begin
    accept Start(ptr : PhilosopherPtr) do
      philo := ptrl;
    end Start;
    work := True;
    while work loop
      Put_Line("Philosopher " & Integer'Image(philo.id) & " is thinking");
      delay 1.0;
      Put_Line("Philosopher " & Integer'Image(philo.id) & " wants to eat");
      philo.leftFork.mtx.lock();
      philo.righFork.mtx.lock();
      Put_Line("Philosopher " & Integer'Image(philo.id) & " is eating");
      delay 1.0;
      philo.righFork.mtx.unlock();
      philo.leftFork.mtx.unlock();
      philo.mtx.lock();
      philo.eatenMeals := philo.eatenMeals + 1;
      work := philo.work;
      philo.mtx.unlock();
    end loop;
  end RunPhilo
end Philosopher;