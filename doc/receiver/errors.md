# OpenTSDB errors

This page should list all the errors OpenTSDB can raise when receiving data.

## HTTP API

### Valid HTTP Request (object)

#### Request

```httpie
http POST ":4243/api/put" metric="test" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 204 No Content
Content-Length: 0
Content-Type: application/json; charset=UTF-8
```

### Valid HTTP Request (array)

#### Request

```httpie
echo "[{\"metric\":\"test1\",\"timestamp\":12345,\"value\":0.1,\"tags\":{\"test1\":\"test1\"}}]" | http POST :4243/api/put 
```

#### Response

```http
HTTP/1.1 204 No Content
Content-Length: 0
Content-Type: application/json; charset=UTF-8
```

### Valid HTTP Request with summary (object)

#### Request

```httpie
http POST ":4243/api/put?summary=true" metric="test" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 50
Content-Type: application/json; charset=UTF-8

{
    "failed": 0,
    "success": 1
}
```

### Valid HTTP Request with summary (array)

#### Request

```httpie
echo "[{\"metric\":\"test1\",\"timestamp\":12345,\"value\":0.1,\"tags\":{\"test1\":\"test1\"}}]" | http POST ":4243/api/put?summary=true"
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 50
Content-Type: application/json; charset=UTF-8

{
    "failed": 0,
    "success": 1
}
```

### Valid HTTP Request with details (object)

#### Request

```httpie
http POST ":4243/api/put?details=true" metric="test" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 61
Content-Type: application/json; charset=UTF-8

{
    "errors": [],
    "failed": 0,
    "success": 1
}
```

### Valid HTTP Request with details (array)

#### Request

```httpie
echo "[{\"metric\":\"test1\",\"timestamp\":12345,\"value\":0.1,\"tags\":{\"test1\":\"test1\"}}]" | http POST ":4243/api/put?details=true"
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 61
Content-Type: application/json; charset=UTF-8

{
    "errors": [],
    "failed": 0,
    "success": 1
}
```

### Invalid HTTP Endpoint

#### Request

```httpie
http :4243/test
```

#### Response

```http
HTTP/1.1 404 Not Found
Content-Encoding: gzip
Content-Length: 481
Content-Type: text/html; charset=UTF-8

<!DOCTYPE html><html><head><meta http-equiv=content-type content="text/html;charset=utf-8"><title>Page Not Found</title>
<style><!--
body{font-family:arial,sans-serif;margin-left:2em}A.l:link{color:#6f6f6f}A.u:link{color:green}.fwf{font-family:monospace;white-space:pre-wrap}//--></style></head>
<body text=#000000 bgcolor=#ffffff><table border=0 cellpadding=2 cellspacing=0 width=100%><tr><td rowspan=3 width=1% nowrap><img src=s/opentsdb_header.jpg><td>&nbsp;</td></tr><tr><td><font color=#507e9b><b>Error 404</b></td></tr><tr><td>&nbsp;</td></tr></table><blockquote><h1>Page Not Found</h1>The requested URL was not found on this server.</blockquote><table width=100% cellpadding=0 cellspacing=0><tr><td class=subg><img alt="" width=1 height=6></td></tr></table></body></html>
```

### Invalid HTTP verb

#### Request

```httpie
http GET :4243/api/put 
```

#### Response

```http
HTTP/1.1 405 Method Not Allowed
Content-Encoding: gzip
Content-Length: 880
Content-Type: application/json; charset=UTF-8

{
    "error": {
        "code": 405,
        "details": "The HTTP method [GET] is not permitted for this endpoint",
        "message": "Method not allowed",
        "trace": "net.opentsdb.tsd.BadRequestException: Method not allowed\n\tat net.opentsdb.tsd.PutDataPointRpc.execute(PutDataPointRpc.java:279) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.handleHttpQuery(RpcHandler.java:282) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.messageReceived(RpcHandler.java:133) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateAwareChannelUpstreamHandler.handleUpstream(IdleStateAwareChannelUpstreamHandler.java:36) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateHandler.messageReceived(IdleStateHandler.java:294) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentEncoder.messageReceived(HttpContentEncoder.java:82) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentDecoder.messageReceived(HttpContentDecoder.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:459) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.callDecode(ReplayingDecoder.java:536) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.messageReceived(ReplayingDecoder.java:435) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:462) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.callDecode(FrameDecoder.java:443) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.messageReceived(FrameDecoder.java:303) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.messageReceived(SimpleChannelHandler.java:142) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat net.opentsdb.tsd.ConnectionManager.handleUpstream(ConnectionManager.java:128) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:559) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:268) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:255) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.read(NioWorker.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.process(AbstractNioWorker.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioSelector.run(AbstractNioSelector.java:337) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.run(AbstractNioWorker.java:89) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.run(NioWorker.java:178) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.ThreadRenamingRunnable.run(ThreadRenamingRunnable.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.internal.DeadLockProofWorker$1.run(DeadLockProofWorker.java:42) [netty-3.10.6.Final.jar:na]\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149) [na:1.8.0_252]\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624) [na:1.8.0_252]\n\tat java.lang.Thread.run(Thread.java:748) [na:1.8.0_252]\n"
    }
}
```

### Unexpected JSON input (object or array)

#### Request

```httpie
http POST :4243/api/put --raw "blah"
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 884
Content-Type: application/json; charset=UTF-8

{
    "error": {
        "code": 400,
        "message": "The JSON must start as an object or an array",
        "trace": "net.opentsdb.tsd.BadRequestException: The JSON must start as an object or an array\n\tat net.opentsdb.tsd.HttpJsonSerializer.parsePutV1(HttpJsonSerializer.java:189) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.PutDataPointRpc.execute(PutDataPointRpc.java:288) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.handleHttpQuery(RpcHandler.java:282) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.messageReceived(RpcHandler.java:133) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateAwareChannelUpstreamHandler.handleUpstream(IdleStateAwareChannelUpstreamHandler.java:36) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateHandler.messageReceived(IdleStateHandler.java:294) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentEncoder.messageReceived(HttpContentEncoder.java:82) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentDecoder.messageReceived(HttpContentDecoder.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:459) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.callDecode(ReplayingDecoder.java:536) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.messageReceived(ReplayingDecoder.java:435) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:462) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.callDecode(FrameDecoder.java:443) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.messageReceived(FrameDecoder.java:303) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.messageReceived(SimpleChannelHandler.java:142) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat net.opentsdb.tsd.ConnectionManager.handleUpstream(ConnectionManager.java:128) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:559) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:268) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:255) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.read(NioWorker.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.process(AbstractNioWorker.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioSelector.run(AbstractNioSelector.java:337) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.run(AbstractNioWorker.java:89) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.run(NioWorker.java:178) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.ThreadRenamingRunnable.run(ThreadRenamingRunnable.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.internal.DeadLockProofWorker$1.run(DeadLockProofWorker.java:42) [netty-3.10.6.Final.jar:na]\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149) [na:1.8.0_252]\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624) [na:1.8.0_252]\n\tat java.lang.Thread.run(Thread.java:748) [na:1.8.0_252]\n"
    }
}
```

### Invalid Content-Type

Not handled by OpenTSDB: it always processes body as JSON.

#### Request

```httpie
 http POST ":4243/api/put?details" Content-Type:application/xml metric="test" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 61
Content-Type: application/json; charset=UTF-8

{
    "errors": [],
    "failed": 0,
    "success": 1
}
```

### Invalid datapoint content (without details)

This error is generic and does not give much details.

#### Request

```httpie
http POST ":4243/api/put" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 911
Content-Type: application/json; charset=UTF-8

{
    "error": {
        "code": 400,
        "details": "Please see the TSD logs or append \"details\" to the put request",
        "message": "One or more data points had errors",
        "trace": "net.opentsdb.tsd.BadRequestException: One or more data points had errors\n\tat net.opentsdb.tsd.PutDataPointRpc$1GroupCB.call(PutDataPointRpc.java:647) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.PutDataPointRpc.processDataPoint(PutDataPointRpc.java:710) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.PutDataPointRpc.execute(PutDataPointRpc.java:296) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.handleHttpQuery(RpcHandler.java:282) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.messageReceived(RpcHandler.java:133) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateAwareChannelUpstreamHandler.handleUpstream(IdleStateAwareChannelUpstreamHandler.java:36) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateHandler.messageReceived(IdleStateHandler.java:294) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentEncoder.messageReceived(HttpContentEncoder.java:82) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentDecoder.messageReceived(HttpContentDecoder.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:459) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.callDecode(ReplayingDecoder.java:536) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.messageReceived(ReplayingDecoder.java:435) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:462) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.callDecode(FrameDecoder.java:443) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.messageReceived(FrameDecoder.java:303) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.messageReceived(SimpleChannelHandler.java:142) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat net.opentsdb.tsd.ConnectionManager.handleUpstream(ConnectionManager.java:128) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:559) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:268) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:255) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.read(NioWorker.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.process(AbstractNioWorker.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioSelector.run(AbstractNioSelector.java:337) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.run(AbstractNioWorker.java:89) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.run(NioWorker.java:178) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.ThreadRenamingRunnable.run(ThreadRenamingRunnable.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.internal.DeadLockProofWorker$1.run(DeadLockProofWorker.java:42) [netty-3.10.6.Final.jar:na]\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149) [na:1.8.0_252]\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624) [na:1.8.0_252]\n\tat java.lang.Thread.run(Thread.java:748) [na:1.8.0_252]\n"
    }
}
```

### Invalid datapoint content (with summary)

#### Request

```httpie
http POST ":4243/api/put?summary=true" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 50
Content-Type: application/json; charset=UTF-8

{
"failed": 1,
"success": 0
}
```

### [DETAILS] Invalid JSON: missing datapoint name

#### Request

```httpie
http POST ":4243/api/put?details=true" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 140
Content-Type: application/json; charset=UTF-8

{
    "errors": [
        {
            "datapoint": {
                "tags": {
                    "test": "test"
                },
                "timestamp": 123456,
                "value": "6.7"
            },
            "error": "Metric name was empty"
        }
    ],
    "failed": 1,
    "success": 0
}
```

### [DETAILS] Invalid JSON: metric name has illegal character

#### Request

```httpie
http POST ":4243/api/put?details=true" metric="Abin76+-" timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 173
Content-Type: application/json; charset=UTF-8

{
    "errors": [
        {
            "datapoint": {
                "metric": "Abin76+-",
                "tags": {
                    "test": "test"
                },
                "timestamp": 123456,
                "value": "6.7"
            },
            "error": "Invalid metric name (\"Abin76+-\"): illegal character: +"
        }
    ],
    "failed": 1,
    "success": 0
}
```

### [DETAILS] Invalid JSON: metric name has float type

#### Request

```httpie
http POST ":4243/api/put?details=true" metric:=7.0 timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 61
Content-Type: application/json; charset=UTF-8

{
    "errors": [],
    "failed": 0,
    "success": 1
}
```

### [DETAILS] Invalid JSON: metric name has map type

In that case, we may have troubled the OpenTSDB JSON parser. It may be handled in our side with the illegal character error.

#### Request

```httpie
http POST ":4243/api/put?details=true" metric:='{"test":"test"}' timestamp:=123456 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 1397
Content-Type: application/json; charset=UTF-8

{
    "error": {
        "code": 400,
        "details": "com.fasterxml.jackson.databind.exc.MismatchedInputException: Cannot deserialize instance of `java.lang.String` out of START_OBJECT token\n at [Source: (String)\"{\"metric\": {\"test\": \"test\"}, \"timestamp\": 123456, \"value\": 6.7, \"tags\": {\"test\": \"test\"}}\"; line: 1, column: 12] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"metric\"])",
        "message": "Unable to parse the given JSON",
        "trace": "net.opentsdb.tsd.BadRequestException: Unable to parse the given JSON\n\tat net.opentsdb.tsd.HttpJsonSerializer.parsePutV1(HttpJsonSerializer.java:192) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.PutDataPointRpc.execute(PutDataPointRpc.java:288) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.handleHttpQuery(RpcHandler.java:282) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.messageReceived(RpcHandler.java:133) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateAwareChannelUpstreamHandler.handleUpstream(IdleStateAwareChannelUpstreamHandler.java:36) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateHandler.messageReceived(IdleStateHandler.java:294) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentEncoder.messageReceived(HttpContentEncoder.java:82) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentDecoder.messageReceived(HttpContentDecoder.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:459) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.callDecode(ReplayingDecoder.java:536) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.messageReceived(ReplayingDecoder.java:435) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:462) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.callDecode(FrameDecoder.java:443) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.messageReceived(FrameDecoder.java:303) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.messageReceived(SimpleChannelHandler.java:142) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat net.opentsdb.tsd.ConnectionManager.handleUpstream(ConnectionManager.java:128) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:559) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:268) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:255) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.read(NioWorker.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.process(AbstractNioWorker.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioSelector.run(AbstractNioSelector.java:337) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.run(AbstractNioWorker.java:89) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.run(NioWorker.java:178) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.ThreadRenamingRunnable.run(ThreadRenamingRunnable.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.internal.DeadLockProofWorker$1.run(DeadLockProofWorker.java:42) [netty-3.10.6.Final.jar:na]\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149) [na:1.8.0_252]\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624) [na:1.8.0_252]\n\tat java.lang.Thread.run(Thread.java:748) [na:1.8.0_252]\nCaused by: java.lang.IllegalArgumentException: com.fasterxml.jackson.databind.exc.MismatchedInputException: Cannot deserialize instance of `java.lang.String` out of START_OBJECT token\n at [Source: (String)\"{\"metric\": {\"test\": \"test\"}, \"timestamp\": 123456, \"value\": 6.7, \"tags\": {\"test\": \"test\"}}\"; line: 1, column: 12] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"metric\"])\n\tat net.opentsdb.utils.JSON.parseToObject(JSON.java:112) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.HttpJsonSerializer.parsePutV1(HttpJsonSerializer.java:181) ~[tsdb-2.4.0.jar:]\n\t... 50 common frames omitted\nCaused by: com.fasterxml.jackson.databind.exc.MismatchedInputException: Cannot deserialize instance of `java.lang.String` out of START_OBJECT token\n at [Source: (String)\"{\"metric\": {\"test\": \"test\"}, \"timestamp\": 123456, \"value\": 6.7, \"tags\": {\"test\": \"test\"}}\"; line: 1, column: 12] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"metric\"])\n\tat com.fasterxml.jackson.databind.exc.MismatchedInputException.from(MismatchedInputException.java:63) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.reportInputMismatch(DeserializationContext.java:1342) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.handleUnexpectedToken(DeserializationContext.java:1138) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.handleUnexpectedToken(DeserializationContext.java:1092) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.StringDeserializer.deserialize(StringDeserializer.java:63) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.StringDeserializer.deserialize(StringDeserializer.java:10) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.impl.MethodProperty.deserializeAndSet(MethodProperty.java:127) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.BeanDeserializer.vanillaDeserialize(BeanDeserializer.java:288) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.BeanDeserializer.deserialize(BeanDeserializer.java:151) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.ObjectMapper._readMapAndClose(ObjectMapper.java:4001) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.ObjectMapper.readValue(ObjectMapper.java:2992) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat net.opentsdb.utils.JSON.parseToObject(JSON.java:108) ~[tsdb-2.4.0.jar:]\n\t... 51 common frames omitted\n"
    }
}
```

### [DETAILS] Invalid datapoint: timestamp is missing

#### Request

```httpie
http POST ":4243/api/put?details=true" metric='test'  value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 135
Content-Type: application/json; charset=UTF-8

{
    "errors": [
        {
            "datapoint": {
                "metric": "test",
                "tags": {
                    "test": "test"
                },
                "timestamp": 0,
                "value": "6.7"
            },
            "error": "Invalid timestamp"
        }
    ],
    "failed": 1,
    "success": 0
}
```

### [DETAILS] Invalid datapoint: timestamp is 0

Timestamp should be > 0.

#### Request

```httpie
http POST ":4243/api/put?details=true" metric='test' timestamp:=0 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 135
Content-Type: application/json; charset=UTF-8

{
    "errors": [
        {
            "datapoint": {
                "metric": "test",
                "tags": {
                    "test": "test"
                },
                "timestamp": 0,
                "value": "6.7"
            },
            "error": "Invalid timestamp"
        }
    ],
    "failed": 1,
    "success": 0
}
```

### [DETAILS] Invalid datapoint: timestamp is float

This is a case OPENTSDB seems to be able to work with…

#### Request

```httpie
http POST ":4243/api/put?details=true" metric='test' timestamp:=1234567.7 value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 61
Content-Type: application/json; charset=UTF-8

{
    "errors": [],
    "failed": 0,
    "success": 1
}
```

### [DETAILS] Invalid datapoint: timestamp is number of format string

The number parser seems to be able to handle number of format string

#### Request

```httpie
 http POST ":4243/api/put?details=true" metric='test' timestamp:="1234567" value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 61
Content-Type: application/json; charset=UTF-8

{
    "errors": [],
    "failed": 0,
    "success": 1
}
```

### [DETAILS] Invalid datapoint: timestamp is string

This cause a general error.

#### Request

```httpie
http POST ":4243/api/put?details=true" metric='test' timestamp="test" value:=6.7 tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 1404
Content-Type: application/json; charset=UTF-8

{
    "error": {
        "code": 400,
        "details": "com.fasterxml.jackson.databind.exc.InvalidFormatException: Cannot deserialize value of type `long` from String \"test\": not a valid Long value\n at [Source: (String)\"{\"metric\": \"test\", \"timestamp\": \"test\", \"value\": 6.7, \"tags\": {\"test\": \"test\"}}\"; line: 1, column: 33] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"timestamp\"])",
        "message": "Unable to parse the given JSON",
        "trace": "net.opentsdb.tsd.BadRequestException: Unable to parse the given JSON\n\tat net.opentsdb.tsd.HttpJsonSerializer.parsePutV1(HttpJsonSerializer.java:192) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.PutDataPointRpc.execute(PutDataPointRpc.java:288) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.handleHttpQuery(RpcHandler.java:282) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.messageReceived(RpcHandler.java:133) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateAwareChannelUpstreamHandler.handleUpstream(IdleStateAwareChannelUpstreamHandler.java:36) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateHandler.messageReceived(IdleStateHandler.java:294) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentEncoder.messageReceived(HttpContentEncoder.java:82) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentDecoder.messageReceived(HttpContentDecoder.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:459) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.callDecode(ReplayingDecoder.java:536) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.messageReceived(ReplayingDecoder.java:435) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:462) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.callDecode(FrameDecoder.java:443) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.messageReceived(FrameDecoder.java:303) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.messageReceived(SimpleChannelHandler.java:142) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat net.opentsdb.tsd.ConnectionManager.handleUpstream(ConnectionManager.java:128) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:559) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:268) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:255) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.read(NioWorker.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.process(AbstractNioWorker.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioSelector.run(AbstractNioSelector.java:337) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.run(AbstractNioWorker.java:89) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.run(NioWorker.java:178) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.ThreadRenamingRunnable.run(ThreadRenamingRunnable.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.internal.DeadLockProofWorker$1.run(DeadLockProofWorker.java:42) [netty-3.10.6.Final.jar:na]\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149) [na:1.8.0_252]\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624) [na:1.8.0_252]\n\tat java.lang.Thread.run(Thread.java:748) [na:1.8.0_252]\nCaused by: java.lang.IllegalArgumentException: com.fasterxml.jackson.databind.exc.InvalidFormatException: Cannot deserialize value of type `long` from String \"test\": not a valid Long value\n at [Source: (String)\"{\"metric\": \"test\", \"timestamp\": \"test\", \"value\": 6.7, \"tags\": {\"test\": \"test\"}}\"; line: 1, column: 33] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"timestamp\"])\n\tat net.opentsdb.utils.JSON.parseToObject(JSON.java:112) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.HttpJsonSerializer.parsePutV1(HttpJsonSerializer.java:181) ~[tsdb-2.4.0.jar:]\n\t... 50 common frames omitted\nCaused by: com.fasterxml.jackson.databind.exc.InvalidFormatException: Cannot deserialize value of type `long` from String \"test\": not a valid Long value\n at [Source: (String)\"{\"metric\": \"test\", \"timestamp\": \"test\", \"value\": 6.7, \"tags\": {\"test\": \"test\"}}\"; line: 1, column: 33] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"timestamp\"])\n\tat com.fasterxml.jackson.databind.exc.InvalidFormatException.from(InvalidFormatException.java:67) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.weirdStringException(DeserializationContext.java:1548) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.handleWeirdStringValue(DeserializationContext.java:910) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.NumberDeserializers$LongDeserializer._parseLong(NumberDeserializers.java:584) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.NumberDeserializers$LongDeserializer.deserialize(NumberDeserializers.java:557) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.NumberDeserializers$LongDeserializer.deserialize(NumberDeserializers.java:535) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.impl.MethodProperty.deserializeAndSet(MethodProperty.java:127) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.BeanDeserializer.vanillaDeserialize(BeanDeserializer.java:288) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.BeanDeserializer.deserialize(BeanDeserializer.java:151) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.ObjectMapper._readMapAndClose(ObjectMapper.java:4001) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.ObjectMapper.readValue(ObjectMapper.java:2992) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat net.opentsdb.utils.JSON.parseToObject(JSON.java:108) ~[tsdb-2.4.0.jar:]\n\t... 51 common frames omitted\n"
    }
}
```

### [DETAILS] Invalid datapoint: value is string

Unlike the timestamp case, an invalid string for value does not cause so much trouble…

#### Request

```httpie
http POST ":4243/api/put?details=true" metric='test' timestamp="123456" value="test" tags:='{"test":"test"}'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 147
Content-Type: application/json; charset=UTF-8

{
    "errors": [
        {
            "datapoint": {
                "metric": "test",
                "tags": {
                    "test": "test"
                },
                "timestamp": 123456,
                "value": "test"
            },
            "error": "Unable to parse value to a number"
        }
    ],
    "failed": 1,
    "success": 0
}
```

### [DETAILS] Invalid datapoint: tags is not a map

#### Request

```httpie
http POST ":4243/api/put?details=true" metric='test' timestamp="123456" value="17778888888" tags:='["test1", "test2"]'
```

#### Response

```http
HTTP/1.1 400 Bad Request
Content-Encoding: gzip
Content-Length: 1437
Content-Type: application/json; charset=UTF-8

{
    "error": {
        "code": 400,
        "details": "com.fasterxml.jackson.databind.exc.MismatchedInputException: Cannot deserialize instance of `java.util.HashMap` out of START_ARRAY token\n at [Source: (String)\"{\"metric\": \"test\", \"timestamp\": \"123456\", \"value\": \"17778888888\", \"tags\": [\"test1\", \"test2\"]}\"; line: 1, column: 75] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"tags\"])",
        "message": "Unable to parse the given JSON",
        "trace": "net.opentsdb.tsd.BadRequestException: Unable to parse the given JSON\n\tat net.opentsdb.tsd.HttpJsonSerializer.parsePutV1(HttpJsonSerializer.java:192) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.PutDataPointRpc.execute(PutDataPointRpc.java:288) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.handleHttpQuery(RpcHandler.java:282) [tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.RpcHandler.messageReceived(RpcHandler.java:133) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateAwareChannelUpstreamHandler.handleUpstream(IdleStateAwareChannelUpstreamHandler.java:36) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.timeout.IdleStateHandler.messageReceived(IdleStateHandler.java:294) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentEncoder.messageReceived(HttpContentEncoder.java:82) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.http.HttpContentDecoder.messageReceived(HttpContentDecoder.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:459) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.callDecode(ReplayingDecoder.java:536) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.replay.ReplayingDecoder.messageReceived(ReplayingDecoder.java:435) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:296) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.unfoldAndFireMessageReceived(FrameDecoder.java:462) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.callDecode(FrameDecoder.java:443) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.handler.codec.frame.FrameDecoder.messageReceived(FrameDecoder.java:303) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelUpstreamHandler.handleUpstream(SimpleChannelUpstreamHandler.java:70) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline$DefaultChannelHandlerContext.sendUpstream(DefaultChannelPipeline.java:791) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.messageReceived(SimpleChannelHandler.java:142) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.SimpleChannelHandler.handleUpstream(SimpleChannelHandler.java:88) [netty-3.10.6.Final.jar:na]\n\tat net.opentsdb.tsd.ConnectionManager.handleUpstream(ConnectionManager.java:128) [tsdb-2.4.0.jar:]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:564) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.DefaultChannelPipeline.sendUpstream(DefaultChannelPipeline.java:559) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:268) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.Channels.fireMessageReceived(Channels.java:255) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.read(NioWorker.java:88) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.process(AbstractNioWorker.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioSelector.run(AbstractNioSelector.java:337) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.AbstractNioWorker.run(AbstractNioWorker.java:89) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.channel.socket.nio.NioWorker.run(NioWorker.java:178) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.ThreadRenamingRunnable.run(ThreadRenamingRunnable.java:108) [netty-3.10.6.Final.jar:na]\n\tat org.jboss.netty.util.internal.DeadLockProofWorker$1.run(DeadLockProofWorker.java:42) [netty-3.10.6.Final.jar:na]\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149) [na:1.8.0_252]\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624) [na:1.8.0_252]\n\tat java.lang.Thread.run(Thread.java:748) [na:1.8.0_252]\nCaused by: java.lang.IllegalArgumentException: com.fasterxml.jackson.databind.exc.MismatchedInputException: Cannot deserialize instance of `java.util.HashMap` out of START_ARRAY token\n at [Source: (String)\"{\"metric\": \"test\", \"timestamp\": \"123456\", \"value\": \"17778888888\", \"tags\": [\"test1\", \"test2\"]}\"; line: 1, column: 75] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"tags\"])\n\tat net.opentsdb.utils.JSON.parseToObject(JSON.java:112) ~[tsdb-2.4.0.jar:]\n\tat net.opentsdb.tsd.HttpJsonSerializer.parsePutV1(HttpJsonSerializer.java:181) ~[tsdb-2.4.0.jar:]\n\t... 50 common frames omitted\nCaused by: com.fasterxml.jackson.databind.exc.MismatchedInputException: Cannot deserialize instance of `java.util.HashMap` out of START_ARRAY token\n at [Source: (String)\"{\"metric\": \"test\", \"timestamp\": \"123456\", \"value\": \"17778888888\", \"tags\": [\"test1\", \"test2\"]}\"; line: 1, column: 75] (through reference chain: net.opentsdb.core.IncomingDataPoint[\"tags\"])\n\tat com.fasterxml.jackson.databind.exc.MismatchedInputException.from(MismatchedInputException.java:63) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.reportInputMismatch(DeserializationContext.java:1342) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.handleUnexpectedToken(DeserializationContext.java:1138) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.DeserializationContext.handleUnexpectedToken(DeserializationContext.java:1092) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.StdDeserializer._deserializeFromEmpty(StdDeserializer.java:599) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.MapDeserializer.deserialize(MapDeserializer.java:360) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.std.MapDeserializer.deserialize(MapDeserializer.java:29) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.impl.MethodProperty.deserializeAndSet(MethodProperty.java:127) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.BeanDeserializer.vanillaDeserialize(BeanDeserializer.java:288) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.deser.BeanDeserializer.deserialize(BeanDeserializer.java:151) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.ObjectMapper._readMapAndClose(ObjectMapper.java:4001) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat com.fasterxml.jackson.databind.ObjectMapper.readValue(ObjectMapper.java:2992) ~[jackson-databind-2.9.5.jar:2.9.5]\n\tat net.opentsdb.utils.JSON.parseToObject(JSON.java:108) ~[tsdb-2.4.0.jar:]\n\t... 51 common frames omitted\n"
    }
}
```
