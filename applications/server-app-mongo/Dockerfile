
FROM openjdk:11-stretch

EXPOSE 8080
EXPOSE 8090

WORKDIR /opt
COPY target/server-app-mongo.jar app.jar

USER 1001

ENV JVM_OPTIMIZATIONS='-noverify -server -XX:TieredStopAtLevel=1'

ENTRYPOINT exec java $JVM_OPTIMIZATIONS -Xms$HEAP_SIZE -Xmx$HEAP_SIZE -XX:MaxMetaspaceSize=$META_SIZE $JAVA_OPTS -jar app.jar
