using System;
using System.Net.Sockets;
using System.Text;
using System.Net;
using System.Collections.Generic;
using System.Threading;

namespace DataLogger
{
	public interface IDataCollector{
		void AddArtifact(string id, IPAddress ip, int port);
	}

	public class DataCollector : IDataCollector
	{
		Dictionary<string, IPEndPoint> artifacts = new Dictionary<string, IPEndPoint>();
		List<string> artifactsToDel = new List<string> ();

		bool running = true;

		public DataCollector ()
		{

		}

		public void AddArtifact(string id, IPAddress ip, int port){
			lock (artifacts) {
				artifacts[id] = new IPEndPoint(ip, port);
			}
		}

		public void GetDataFromAllArtifacts ()
		{
			lock (artifacts) {
				foreach (var artifact in artifacts.Keys) {
					try {
						string data = GetData (artifacts[artifact]);
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
				Thread.Sleep(1000);
				GetDataFromAllArtifacts();   
			};
		}

		public void Start(){
			Thread t = new Thread(new ThreadStart(this.DataCollectorMain));
			t.Start();
		}

		string GetData (IPEndPoint artifact)
		{
			Socket server = new Socket (AddressFamily.InterNetwork,
		                     SocketType.Stream, ProtocolType.Tcp);
			server.Connect (artifact);
			server.Send (Encoding.ASCII.GetBytes ("action=request request=all"));
			byte[] data = new byte[1024];
            int length = server.Receive(data);
            server.Close();
			return Encoding.ASCII.GetString (data,0,length);
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

