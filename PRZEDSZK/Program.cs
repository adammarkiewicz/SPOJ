using System;

namespace PRZEDSZK
{
    class Program
    {
        static void Main(string[] args)
        {
            byte testsCount = Convert.ToByte(Console.ReadLine());

            while (testsCount-- > 0)
            {
                string[] inputs = Console.ReadLine().Split(' ');

                byte g1 = Convert.ToByte(inputs[0]);
                byte g2 = Convert.ToByte(inputs[1]);

                if (g1 < g2)
                {
                    var tmp = g1;
                    g1 = g2;
                    g2 = tmp;
                }

                for (int i = 1; i <= g2; i++)
                {
                    if (g1 * i % g2 == 0)
                    {
                        Console.WriteLine(i * g1);
                        break;
                    }
                }
            }
        }
    }
}
