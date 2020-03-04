# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import realtime_service_pb2 as realtime__service__pb2


class RealtimeServiceStub(object):
  """Realtime service def
  Remote Procedure Calls
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetLatestStreaming = channel.unary_stream(
        '/realtime.RealtimeService/GetLatestStreaming',
        request_serializer=realtime__service__pb2.RouteRequest.SerializeToString,
        response_deserializer=realtime__service__pb2.RouteResponse.FromString,
        )
    self.GetLatest = channel.unary_stream(
        '/realtime.RealtimeService/GetLatest',
        request_serializer=realtime__service__pb2.RouteRequest.SerializeToString,
        response_deserializer=realtime__service__pb2.RouteResponse.FromString,
        )


class RealtimeServiceServicer(object):
  """Realtime service def
  Remote Procedure Calls
  """

  def GetLatestStreaming(self, request, context):
    """
    GetLatestStreaming
    returns latest values, and future values for provided namespace
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetLatest(self, request, context):
    """
    GetLatest
    returns latest values for provided namespace
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RealtimeServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetLatestStreaming': grpc.unary_stream_rpc_method_handler(
          servicer.GetLatestStreaming,
          request_deserializer=realtime__service__pb2.RouteRequest.FromString,
          response_serializer=realtime__service__pb2.RouteResponse.SerializeToString,
      ),
      'GetLatest': grpc.unary_stream_rpc_method_handler(
          servicer.GetLatest,
          request_deserializer=realtime__service__pb2.RouteRequest.FromString,
          response_serializer=realtime__service__pb2.RouteResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'realtime.RealtimeService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
