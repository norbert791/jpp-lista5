with Ada.Text_IO;
use Ada.Text_IO;
with Philosopher;
use Philosopher;

procedure main is
  fork1 : ForkPtr;
  fork2 : ForkPtr;
  philo1 : PhilosopherPtr;
  philo2 : PhilosopherPtr;
  run1 : RunPhilo;
  run2 : RunPhilo;
  work : Boolean;
begin
  fork1 := new Fork;
  fork1.id := 0;
  --fork1.mtx := MyMutex;
  fork2 := new Fork;
  fork2.id := 1;
  --fork2.mtx := MyMutex;
  philo1 := new philosopher.Philosopher;
  philo2 := new philosopher.Philosopher;
  philo1.eatenMeals := 0;
  philo1.id := 0;
  philo1.leftFork := fork1;
  philo1.righFork := fork2;
  --philo1.mtx := MyMutex;
  philo1.work := True;
  philo2.eatenMeals := 0;
  philo2.id := 1;
  philo2.leftFork := fork2;
  philo2.righFork := fork1;
  --philo2.mtx := MyMutex;
  philo2.work := True;

  work := True;
  run1.Start(ptr => philo1);
  run2.Start(ptr => philo2);

  while work loop
    delay 1.0;
    work := False;
    philo1.mtx.lock;
    if philo1.eatenMeals < 5 then
      work := True;
    end if;
    philo1.mtx.unlock;
    philo2.mtx.lock;
    if philo2.eatenMeals < 5 then
      work := True;
    end if;
    philo2.mtx.unlock;
  end loop;

  philo1.mtx.lock;
  philo1.work := False;
  philo1.mtx.unlock;
  philo2.mtx.lock;
  philo2.work := False;
  philo2.mtx.unlock;

end main;