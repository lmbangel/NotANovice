import grpc
from concurrent import futures
import quiz_pb2
import quiz_pb2_grpc
import ollama
import json
# from ollama import chat
# from ollama import ChatResponse


class QuizAgentService(quiz_pb2_grpc.QuizAgentServiceServicer):
    def GetQuizOfTheDay(self, request, context):
        with open("_prompt.md", "r", encoding="utf-8") as f:
            data = f.read()
        
        response = ollama.chat(
            model="llama3.2:latest",
            messages=[{"role": "user", "content": data}],
            format={
                "type":"object",
                "properties":{
                    "question": {"type": "string"},
                    "correct_answer": {"type":"string"},
                    "timestamp":{"type": "integer"},
                    "a_answer": {"type":"string"},
                    "b_answer": {"type":"string"},
                    "c_answer": {"type":"string"},
                    "d_answer": {"type":"string"}
                }
            },
            options={
                "temperature": 0
            },
            stream=False
        )
        content = response.message.content
        
        try:
            quiz = json.loads(content)
        except json.JSONDecodeError:
            print("Model did not return valid JSON")
            quiz = {"raw_output": content}
          
        return quiz_pb2.Quiz(
            quiz_id=1,
            id=2,
            question=quiz.get("question", ""),
            correct_answer=quiz.get("correct_answer", ""),
            timestamp=quiz.get("timestamp", 0),
            is_active=1,
            a_answer=quiz.get("a_answer", ""),
            b_answer=quiz.get("b_answer", ""),
            c_answer=quiz.get("c_answer", ""),
            d_answer=quiz.get("d_answer", "")
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    quiz_pb2_grpc.add_QuizAgentServiceServicer_to_server(QuizAgentService(), server)
    server.add_insecure_port("[::]:50052")
    server.start()
    print("Python gRPC server running on :50052")

    server.wait_for_termination()


if __name__ == "__main__":
    serve()
