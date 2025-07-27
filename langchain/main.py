from load import load_documents_from_directory
from splitter import split_documentss
from vector_store import store_to_pgvector
from rag_qa import rag_question
import os
os.environ["http_proxy"] = "http://127.0.0.1:7890"
os.environ["https_proxy"] = "http://127.0.0.1:7890"
# 配置参数
DATA_DIR = "./doc"
PG_CONN = ""
OPENAI_API_KEY = "sk"

if __name__ == "__main__":
    # 1. 加载数据
    docs = load_documents_from_directory(DATA_DIR)
    # 2. 分块
    chunks = split_documentss(docs)
    # print(chunks)
    # 3. 存储到PG向量数据库
    vectorstore = store_to_pgvector(chunks, PG_CONN, OPENAI_API_KEY)
    # 4. 用户提问
    query = input("请输入你的问题：")
    # 5. 检索并问答
    result = rag_question(vectorstore, OPENAI_API_KEY, query)
    print(result)