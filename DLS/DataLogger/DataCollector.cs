using System;
using System.Net.Sockets;
using System.Text;
using System.Net;
using System.Collections.Generic;
using System.Threading;

namespace DataLogger
{
	public class DataCollector
	{
		List<IPEndPoint> artifacts = new List<IPEndPoint> ();
		List<IPEndPoint> artifactsToDel = new List<IPEndPoint> ();
		bool running = true;

		public DataCollector ()
		{

		}

		public void AddArtifact (IPEndPoint ipep)
		{
			lock (artifacts) {
				artifacts.Add (ipep);
			}
		}

		public void AddArtifact(string id, IPAddress ip, int port){
			lock (artifacts) {
				artifacts.Add (new IPEndPoint(ip, port));
			}
		}

		public void GetDataFromAllArtifacts ()
		{
			lock (artifacts) {
				foreach (var artifact in artifacts) {
					try {
						string data = GetData (artifact);
						Console.WriteLine (data);

					} catch (System.Net.Sockets.SocketException) {
						artifactsToDel.Add (artifact);
					}
				}
			}
			CleanArtifacts ();
		}

		public void DataCollectorMain(){
			while(running){
				Thread.Sleep(500);
				GetDataFromAllArtifacts();
			};
		}

		public void Start(){
			Thread t = new Thread(new ThreadStart(this.DataCollectorMain));
			t.Start();
		}

		static string GetData (IPEndPoint artifact)
		{
			Socket server = new Socket (AddressFamily.InterNetwork,
		                     SocketType.Stream, ProtocolType.Tcp);
			server.Connect (artifact);
			server.Send (Encoding.ASCII.GetBytes ("request=all"));
			byte[] data = new byte[1024];
			server.Receive (data);
			string message = Encoding.ASCII.GetString (data);
			return message;
		}

		void CleanArtifacts ()
		{
			lock(artifacts){
			foreach (var artifact in artifactsToDel) {
				artifacts.Remove (artifact);
			}
			}
		}
	}
}

