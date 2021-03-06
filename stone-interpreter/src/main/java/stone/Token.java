package stone;

public class Token {

  public static final Token EOF = new Token(-1) {
  }; // end of file
  public static final String EOL = "\\n"; // end of line
  private int lineNumber;

  protected Token(int line) {
    this.lineNumber = line;
  }

  public boolean isNumber() {
    return false;
  }

  public boolean isString() {
    return false;
  }

  public boolean isIdentifier() {
    return false;
  }

  public int getNumber() {
    throw new StoneException("not number token");
  }

  public int getLineNumber() {
    return lineNumber;
  }

  public String getText() {
    return "";
  }

}
