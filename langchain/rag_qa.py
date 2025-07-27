from langchain.llms import OpenAI
from langchain.prompts import PromptTemplate
from langchain.chains import RetrievalQA
from langchain_community.llms import QwenTurbo

def rag_question(vectorstore, openai_api_key, query, k=5):
    # llm = OpenAI(openai_api_key=openai_api_key)
    llm = QwenTurbo(api_key="sk-96f10ba7e8db41c4a45bf2908401bafb", api_base="https://dashscope.aliyuncs.com/compatible-mode/v1")
    prompt = PromptTemplate(
        template="请根据以下内容回答问题：\n{context}\n问题：{question}",
        input_variables=["context", "question"]
    )
    qa_chain = RetrievalQA.from_chain_type(
        llm=llm,
        chain_type="stuff",
        retriever=vectorstore.as_retriever(search_kwargs={"k": k}),
        return_source_documents=True,
        chain_type_kwargs={"prompt": prompt}
    )
    return qa_chain(query)