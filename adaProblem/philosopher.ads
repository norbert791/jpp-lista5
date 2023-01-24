package philosopher is

  protected type MyMutex is
    entry lock;
    entry unlock;
  private
    locked : Boolean := False; 
  end MyMutex;

  type Fork is record
    id : Integer;
    mtx : MyMutex;
  end record;

  type ForkPtr is access Fork;

  type Philosopher is record
    id : Integer;
    mtx : MyMutex;
    eatenMeals : Integer;
    work : Boolean;
    leftFork : ForkPtr;
    righFork : ForkPtr;
  end record;

  type PhilosopherPtr is access Philosopher;

  task type RunPhilo is
    entry Start(ptr : PhilosopherPtr);
  end RunPhilo;

end Philosopher;