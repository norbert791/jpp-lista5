with Ada.Text_IO;
use Ada.Text_IO;
with Philosopher;
use Philosopher;

procedure main is
  type philoNum is range 0..5;
  type philoArr is array(philoNum) of PhilosopherPtr;
  type forkArr is array(philoNum) of ForkPtr;
  type threadArr is array(philoNum) of RunPhilo;
  philoIdx : philoNum := 0;
  idPool : Integer := 0;
  philos : philoArr;
  forks : forkArr;
  work : Boolean := True;
  maxMeals : Integer := 5;
  threads : threadArr;
begin

  for index in philoNum'Range loop
    philos(index) := new philosopher.Philosopher;
    philos(index).id := idPool;
    philos(index).work := True;
    philos(index).eatenMeals := 0;
    forks(index) := new Fork;
    forks(index).id := idPool;
    idPool := idPool + 1;
  end loop;

  for index in philoNum'Range loop
    philos(index).leftFork := forks(index);
    if index = philoNum'Last then
      philos(index).righFork := forks(0);
    else
      philos(index).righFork := forks(index + 1);
    end if;
  end loop;

  for index in philoNum'Range loop
    threads(index).Start(philos(index));
  end loop;

  while work loop
    delay 1.0;
    work := False;
    for index in philoNum'Range loop
      philos(index).mtx.lock;
      if philos(index).eatenMeals < maxMeals then
        work := True;
      end if;
      philos(index).mtx.unlock;
    end loop;
  end loop;

  for index in philoNum'Range loop
    philos(index).mtx.lock;
    philos(index).work := False;
    philos(index).mtx.unlock;
  end loop;
end main;