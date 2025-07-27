import os
from langchain_community.document_loaders import TextLoader, PyPDFLoader

def load_documents_from_directory(directory):
    docs = []
    for root, _, files in os.walk(directory):
        for file in files:
            print(f"正在加载文件: {file}")
            path = os.path.join(root, file)
            try:
                if file.lower().endswith(".txt"):
                    loader = TextLoader(path, encoding="utf-8")
                    docs.extend(loader.load())
                elif file.lower().endswith(".pdf"):
                    loader = PyPDFLoader(path)
                    docs.extend(loader.load())
            except Exception as e:
                print(f"跳过文件 {file}，原因：{e}")
    return docs

if __name__ == "__main__":
    docs = load_documents_from_directory("./doc")
    for doc in docs:
        print(doc.page_content)