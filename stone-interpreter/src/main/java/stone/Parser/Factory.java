package stone.Parser;

import java.lang.reflect.Constructor;
import java.lang.reflect.Method;
import java.util.List;
import stone.ast.ASTree;

public abstract class Factory {

  public static final String factoryName = "create";

  protected static Factory get(Class<? extends ASTree> clazz, Class<?> argType) {
    if (clazz == null) {
      return null;
    }
    try {
      final Method m = clazz.getMethod(factoryName, new Class<?>[]{argType});
    } catch (NoSuchMethodException e) {

    }

    try {
      final Constructor<? extends ASTree> c = clazz.getConstructor(argType);
    } catch (NoSuchMethodException e) {
      throw new RuntimeException(e);
    }

  }

  protected static Factory getForASTList(Class<? extends ASTree> clazz) {
    Factory f = get(clazz, List.class);
    return f;
  }

  protected abstract ASTree make0(Object arg) throws Exception;

  protected ASTree make(Object arg) throws IllegalArgumentException {
    /**
     * this is basically a wrapper around make0
     * it allows you to throw only illegal argument Exception.
     * Any other issues are thrown as runtime exception and make the compiler to be broken
     *
     * */
    try {
      return make0(arg);
    } catch (IllegalArgumentException e1) {
      throw e1;
    } catch (Exception e2) {
      throw new RuntimeException(e2); // compiler is broken
    }
  }
}
