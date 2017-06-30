package tantanwen.www;

import org.apache.thrift.TException;

/**
 * Created by dan on 2017/6/29.
 */
public class HelloServiceImpl implements Hello.Iface {
    @Override
    public String helloString(String para) throws TException {
        return null;
    }

    @Override
    public int helloInt(int para) throws TException {
        return para;
    }

    @Override
    public boolean helloBoolean(boolean para) throws TException {
        return false;
    }

    @Override
    public void helloVoid() throws TException {
        System.out.println("Hello World");
    }

    @Override
    public String helloNull() throws TException {
        return null;
    }
}
