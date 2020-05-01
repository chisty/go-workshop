using System;
using System.ComponentModel;
using Grpc.Core;
using Proto;


namespace Client
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("GRPC Client. Enter 2 integer separated by space. (Enter blank to exit)");

            var channel = new Channel("localhost:8080", ChannelCredentials.Insecure);
            var client = new CalculationService.CalculationServiceClient(channel);


            while (true)
            {
                var input = Console.ReadLine();
                if (string.IsNullOrEmpty(input)) break;

                var tokens = input.Split(' ');
                long a, b;

                if (tokens.Length == 2 && long.TryParse(tokens[0], out a) && long.TryParse(tokens[1], out b))
                {
                    Add(a, b, client);
                    Multiply(a,b,client);
                }
                else
                {
                    Console.WriteLine("Invalid input: "+input);
                }
            }
        }


        private static void Add(long a, long b, CalculationService.CalculationServiceClient client)
        {
            var request = new Request { A = a, B = b };
            var reply = client.Add(request);
            Console.WriteLine("Result of Addition (From Go Server): " + reply.Result);
        }
        private static void Multiply(long a, long b, CalculationService.CalculationServiceClient client)
        {
            var request = new Request { A = a, B = b };
            var reply = client.Multiply(request);
            Console.WriteLine("Result of Multiplication (From Go Server): " + reply.Result);
        }

    }
    

    //tools\protoc.exe -I proto proto\service.proto --csharp_out=proto --grpc_out=proto --plugin=protoc-gen-grpc=tools\grpc_csharp_plugin.exe
}
