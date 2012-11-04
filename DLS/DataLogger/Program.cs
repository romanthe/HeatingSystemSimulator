using System;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;
using System.Net.Sockets;
using System.Text;
using System.Net;

namespace DataLogger
{
    static class Program
    {
        public static void Main()
        {
			DataCollector dc = new DataCollector();
			dc.Start();
			RegistrationServer rs = new RegistrationServer(dc);
			rs.Start();
		}
    }
}
