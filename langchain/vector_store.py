# from langchain.embeddings import OpenAIEmbeddings
from langchain.vectorstores.pgvector import PGVector
# # 替换为阿里百炼的 Embedding 和 LLM
# from langchain_community.embeddings import QwenEmbeddings
# from langchain_community.llms import QwenTurbo

from langchain_community.embeddings import HuggingFaceEmbeddings
# from langchain_community.vectorstores import PGVector


# llm = QwenTurbo(api_key="你的阿里百炼API_KEY", api_base="https://dashscope.aliyuncs.com/")

def store_to_pgvector(docs, connection_string, openai_api_key, collection_name="rag_collection"):
    # embeddings = OpenAIEmbeddings(openai_api_key=openai_api_key)
    # embeddings = QwenEmbeddings(api_key="sk-96f10ba7e8db41c4a45bf2908401bafb", api_base="https://dashscope.aliyuncs.com/compatible-mode/v1") 
    embeddings = HuggingFaceEmbeddings(model_name="shibing624/text2vec-base-chinese")
    print("生成embedding", embeddings)
    vectorstore = PGVector.from_documents(
        docs,
        embeddings,
        connection_string=connection_string,
        collection_name=collection_name
    )
    return vectorstore

