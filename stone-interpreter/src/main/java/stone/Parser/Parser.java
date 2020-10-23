package stone.Parser;

import java.lang.reflect.Constructor;
import java.lang.reflect.Method;
import java.util.List;
import stone.Lexer;
import stone.ParseException;
import stone.ast.ASTList;
import stone.ast.ASTree;

public class Parser {


  public static final String factoryName = "create";
  protected List<Element> elements;
  protected Factory factory;

  protected static abstract class Factory {

    protected static Factory getForASTList(Class<? extends ASTree> clazz) {
      Factory f = get(clazz, List.class);
      if (f ==  null) {
        f = new Factory() {
          @Override
          protected ASTree make0(Object arg) throws Exception {
            List<ASTree> results =  (List<ASTree>) arg ;
            if (results.size() == 1) {
              return results.get(0);
            }
            return new ASTList(results);
          }
        };
      }
      return f;
    }

    protected static Factory get(Class<? extends ASTree> clazz, Class<?> argType) {
      if (clazz == null) {
        return null;
      }
      try {
        final Method m = clazz.getMethod(factoryName, new Class<?>[]{argType});
        return new Factory() {
          @Override
          protected ASTree make0(Object arg) throws Exception {
            return (ASTree) m.invoke(null, arg);
          }
        };
      } catch (NoSuchMethodException e) {
      }

      try {
        final Constructor<? extends ASTree> c = clazz.getConstructor(argType);
        return new Factory() {
          @Override
          protected ASTree make0(Object arg) throws Exception {
            return c.newInstance(arg);
          }
        };
      } catch (NoSuchMethodException e) {
        throw new RuntimeException(e);
      }


    }

    protected abstract ASTree make0(Object arg) throws Exception;

    protected ASTree make(Object arg) {
      try {
        return make0(arg);
      } catch (IllegalArgumentException e1) {
        throw e1;
      } catch (Exception e2) {
        throw new RuntimeException(e2); // compiler is broken
      }
    }
  }


  protected static abstract class Element {

    protected abstract void parse(Lexer lexer, List<ASTree> res) throws ParseException;

    protected abstract boolean match(Lexer lexer) throws ParseException;
  }

  protected static class Tree extends Element {

    protected Parser parser;

    protected Tree(Parser p) {
      parser = p;
    }


    protected void parse(Lexer lexer, List<ASTree> res) throws ParseException {
      res.add(parser.parse(lexer);
    }

    protected boolean match(Lexer lexer) throws ParseException {
      return false;
    }
  }
}
