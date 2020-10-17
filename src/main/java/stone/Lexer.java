package stone;

import java.io.IOException;
import java.io.LineNumberReader;
import java.io.Reader;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;


public class Lexer {

  private static String SPACE = "\\s*";  // in regular expression, we need to
  private static String COMMENT = "(//.*)";
  private static String NUMBER_LITERAL = "([0-9]+)";
  private static String STRING_LITERAL = "(\"(\\\"|\\\\|\\\\n|[^\"])*\")";

  // start with "
  // there are 4 conditions in the middle:
  // a.  \" => translate as \\\"
  // b.  \
  // c.  \n
  // d. anything not starting with "
  // end with "
  private static String VAR_NAME = "[A-Za-z][A-Za-z0-9)*";
  private static String SPECIAL_OPERATOR = "==|&&|\\|\\||>=|<=|\\p{Punct}";
  private static String regexPat = String
      .format("%s(%s|%s|%s|%s|%s)?", SPACE, COMMENT, NUMBER_LITERAL,
          STRING_LITERAL, VAR_NAME, SPECIAL_OPERATOR);
  private Pattern pattern = Pattern.compile(regexPat);
  private List<Token> queue = new ArrayList<Token>();

  private boolean hasMore;
  private LineNumberReader reader;

  public Lexer(Reader r) {
    this.hasMore = true;
    this.reader = new LineNumberReader(r);
  }

  public Token read() throws ParseException {
    if (fillQueue(0)) {
      return queue.remove(0);
    } else {
      return Token.EOF;
    }
  }

  public Token peek(int i) throws ParseException {
    if (fillQueue(i)) {
      return queue.get(i);
    } else {
      return Token.EOF;
    }
  }

  private boolean fillQueue(int i) throws ParseException {
    while (i >= queue.size()) {
      if (this.hasMore) {
        this.readLine();
      } else {
        return false;
      }
    }
    return true;
  }

  private void readLine() throws ParseException {
    String line;
    try {
      line = this.reader.readLine();
    } catch (IOException e) {
      throw new ParseException(e);
    }
    if (line == null) {
      this.hasMore = false;
      return;
    }
    int lineNo = reader.getLineNumber();
    Matcher matcher = pattern.matcher(line);
    matcher.useAnchoringBounds(false)
        .useTransparentBounds(true);
    // transparent bound -> allows the match to go beyond the bound
    // anchoring bound -> respect ^ and $, the default is  True
    int pos = 0;
    int endPos = line.length();
    while (pos < endPos) {
      matcher.region(pos, endPos);
      if (matcher.lookingAt()) {
        addToken(lineNo, matcher);
        pos = matcher.end();
      } else {
        throw new ParseException("bad token at line " + lineNo);
      }
    }
    queue.add(new IdToken(lineNo, Token.EOL));
  }


  protected void addToken(int lineNo, Matcher matcher) {
    String m = matcher.group(1);
    if (m != null) { // group 0 is space
      if (matcher.group(2) == null) // if not a comment
      {
        Token token;
        if (matcher.group(3) != null) { //is numToken
          token = new NumToken(lineNo, Integer.parseInt(m));
        } else if (matcher.group(4) != null) { // is string token
          token = new StrToken(lineNo, toStringLiteral(m));
        } else { // id token
          token = new IdToken(lineNo, m);
        }
        queue.add(token);
      }
    }
  }

  protected String toStringLiteral(String s) {
    StringBuilder sb = new StringBuilder();
    int len = s.length() - 1;
    for (int i = 1; i < len; i++) { // only case anything within "XX"
      char c = s.charAt(i);
      if (c == '\\' && i + 1 < len) { // determine if it is an escape \
        int c2 = s.charAt(i + 1);
        if (c2 == '"' || c2 == '\\') { //  \" or \\
          c = s.charAt(++i);
        } else if (c2 == 'n') {
          ++i;
          c = '\n';
        }
      }
      sb.append(c);
    }
    return sb.toString();
  }

  protected static class NumToken extends Token {

    private int value;

    protected NumToken(int line, int value) {
      super(line);
      this.value = value;
    }


    @Override
    public boolean isNumber() {
      return true;
    }

    @Override
    public String getText() {
      return Integer.toString(this.value);
    }

    @Override
    public int getNumber() {
      return this.value;
    }
  }

  protected static class IdToken extends Token {

    private String text;

    IdToken(int line, String text) {
      super(line);
      this.text = text;
    }

    @Override
    public boolean isIdentifier() {
      return true;
    }

    @Override
    public String getText() {
      return text;
    }
  }

  protected static class StrToken extends Token {

    private String literal;

    StrToken(int line, String str) {
      super(line);
      this.literal = str;
    }

    @Override
    public boolean isString() {
      return true;
    }

    @Override
    public String getText() {
      return this.literal;
    }
  }
}




