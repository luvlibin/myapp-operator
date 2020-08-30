package abc

var configmapTemplate = `
-Dfile.encoding=UTF8
-Dgw.server.mode=test
-Dgw.cc.serverid=$HOSTNAME#ui,batch,workqueue,scheduler,messaging,startable
-Dgw.cc.env=mytest-operator
-Dadj.cc.saml.properties.file=/usr/local/tomcat/configuration/saml.cc.properties
-Dadj.properties.path=/usr/local/tomcat/configuration/adj.cc.properties
-server
-XX:+UseParallelGC
-Djava.net.preferIPv4Stack=true
-Djava.awt.headless=true
-XX:+UseCompressedOops
-Xms4500m
-Xmx4500m
-XX:MetaspaceSize=512m
-XX:MaxMetaspaceSize=512m
-XX:ReservedCodeCacheSize=80m
-Dcatalina.log.path=/tmp/gwlogs/abc/logs/$HOSTNAME
`
