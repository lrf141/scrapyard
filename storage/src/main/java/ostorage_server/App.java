package ostorage_server;

import org.eclipse.jetty.server.*;
import org.eclipse.jetty.server.handler.HandlerList;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;
import ostorage_server.handler.HelloHandler;

public class App {

    public static void main(String[] args) {
        ServletContextHandler contextHandler = new ServletContextHandler();
        contextHandler.setMaxFormContentSize(1024);
        contextHandler.addServlet(new ServletHolder(new HelloHandler()), "/");

        HandlerList handlerList = new HandlerList();
        handlerList.addHandler(contextHandler);

        final int port = 8000;
        final Server jettyServer = new Server();
        jettyServer.setHandler(handlerList);

        final HttpConfiguration conf = new HttpConfiguration();
        conf.setSendServerVersion(false);

        final HttpConnectionFactory httpConnectionFactory = new HttpConnectionFactory(conf);
        final ServerConnector serverConnector = new ServerConnector(jettyServer, httpConnectionFactory);
        serverConnector.setPort(port);
        jettyServer.setConnectors(new Connector[] { serverConnector });

        try {
            jettyServer.start();
            jettyServer.join();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
