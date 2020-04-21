using System;
using System.Text;

namespace PP0504D
{
    class Program
    {
        static void Main()
        {
            short testsCount = Convert.ToInt16(Console.ReadLine());
            StringBuilder output = new StringBuilder();

            while (testsCount-- > 0)
            {
                float x = float.Parse(Console.ReadLine());
                byte[] bytes = BitConverter.GetBytes(x);
                StringBuilder hex = new StringBuilder(bytes.Length * 2);

                for (int i = bytes.Length - 1; i >= 0; i--)
                {
                    hex.AppendFormat("{0:x} ", bytes[i]);
                }

                output.Append(hex + Environment.NewLine);
            }

            Console.WriteLine(output);
        }
    }
}
