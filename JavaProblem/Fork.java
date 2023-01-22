package JavaProblem;

public class Fork{
  public Fork() {id = generateId();}
  public int getId() {return id;}
  
  final private int id;

  private static volatile int idPool = 0;
  private static synchronized int generateId() {return idPool++;}
}
