import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TSimpleServer;
import org.apache.thrift.transport.TServerSocket;
import org.apache.thrift.transport.TServerTransport;
import tantanwen.www.Hello;
import tantanwen.www.HelloServiceImpl;
import org.apache.thrift.server.TServer.Args;

public class HelloServiceServer {
    public static HelloServiceImpl handler;
    public static Hello.Processor processor;
    public static void main(String[] args) {
        try {
        //注册成方法
        handler = new HelloServiceImpl();
        processor = new Hello.Processor(handler);
        Runnable simple = new Runnable() {
            public void run() {
                simple(processor);
            }
        };
        new Thread(simple).start();
        } catch (Exception x) {
            x.printStackTrace();
        }
        /*
        try {
            // 设置服务端口为 7911
            TServerSocket serverTransport = new TServerSocket(7911);
            // 设置协议工厂为 TBinaryProtocol.Factory
            TBinaryProtocol.Factory proFactory = new TBinaryProtocol.Factory();
            // 关联处理器与 Hello 服务的实现
            TProcessor processor = new Hello.Processor(new HelloServiceImpl());
            //TServer server = new TThreadPoolServer(processor, serverTransport, proFactory);
            TServer server = new TThreadPoolServer(new TThreadPoolServer.Args(serverTransport).processor(processor));
            System.out.println("Start server on port 7911...");
            server.serve();
        } catch (TTransportException e) {
            e.printStackTrace();
        }*/
    }
    public static void simple(Hello.Processor processor) {
        try {
            TServerTransport serverTransport = new TServerSocket(9090);
            TServer server = new TSimpleServer(new Args(serverTransport).processor(processor));

            // Use this for a multithreaded server
            // TServer server = new TThreadPoolServer(new TThreadPoolServer.Args(serverTransport).processor(processor));

            System.out.println("Starting the simple server...");
            server.serve();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}