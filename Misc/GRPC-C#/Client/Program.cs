using System;
using Grpc.Core;
using Service;


namespace Client
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");

            var channel = new Channel("127.0.0.1:5001", ChannelCredentials.Insecure);
            var client = new CalculationService.CalculationServiceClient(channel);

            var request = new Request {A = 2, B = 3};
            var reply = client.Add(request);
            Console.WriteLine("Result from Go Server: "+reply.Result);

            reply = client.Multiply(request);
            Console.WriteLine("Result from Go Server: " + reply.Result);

            Console.ReadKey();
        }
    }



    //tools\protoc.exe -I protos protos\test.proto --csharp_out=output --grpc_out=output --plugin=protoc-gen-grpc=tools\grpc_csharp_plugin.exe
}
