package stone.ast;

import java.util.ArrayList;
import java.util.Iterator;
import stone.Token;

public class ASTLeaf extends ASTree {

  private static ArrayList<ASTree> empty = new ArrayList<ASTree>();

  protected Token token;

  public ASTLeaf(Token token) {
    this.token = token;
  }

  @Override
  public ASTree child(int i) {
    throw new IndexOutOfBoundsException();
  }

  public int numChildren() {
    return 0;
  }

  public Iterator<ASTree> children() {
    return empty.iterator();
  }

  public String location() {
    return "at line " + this.token.getLineNumber();
  }

  public Token token() {
    return this.token;
  }
}
