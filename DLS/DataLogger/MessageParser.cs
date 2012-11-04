using System;
using System.Linq;

namespace DataLogger
{
	public class NoSuchFieldException : Exception{

	}

	public class MessageParser
	{
		public string Message{get; set;}
		public MessageParser (string message)
		{
			Message = message;
		}

		public string GetString(string field){
			string[] pairs = Message.Split(' ');
			var ports = from pair in pairs where pair.Contains(field + "=") select pair.Remove(0,field.Length + 1);
			string answer = ports.FirstOrDefault();
			if(answer != null){
				return answer;
			}else
				throw new NoSuchFieldException();
		}

		public int GetInt(string field){
			return int.Parse(GetString(field));
		}
	}
}

