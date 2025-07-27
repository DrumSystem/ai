from langchain.text_splitter import RecursiveCharacterTextSplitter

def split_documentss(docs, chunk_size=500, chunk_overlap=100):
    splitter = RecursiveCharacterTextSplitter(chunk_size=chunk_size, chunk_overlap=chunk_overlap)
    return splitter.split_documents(docs)