# Prerequisites

## Redpanda Serverless

Log in to your Redpanda Cloud account (or sign up if you don't have one) and create a new [Serverless](https://redpanda.com/redpanda-cloud/serverless) cluster. Make a note of the bootstrap server URL, and create a new user with permissions (ACLs) to access a topic named `documents` and a consumer group named `benthos`. Add the cluster connection information to a local `.env` file:

```bash
% cat > demo/.env<< EOF
REDPANDA_SERVERS="<bootstrap_server_url>"
REDPANDA_USER="<username>"
REDPANDA_PASS="<password>"
REDPANDA_TOPIC="documents"

EOF
```

## OpenAI API

Log in to your OpenAI developer platform account (or sign up if you don't have one) and create a new [Project API key](https://platform.openai.com/api-keys). Add the secret key to the local `.env` file:

```bash
% cat >> demo/.env<< EOF
OPENAI_API_KEY="<secret_key>"
OPENAI_EMBEDDING_MODEL="text-embedding-3-small"
OPENAI_MODEL="gpt-4o"

EOF
```

## MongoDB Atlas

Log in to your MongoDB Atlas account (or sign up if you don't have one) and deploy a new [free cluster](https://www.mongodb.com/docs/atlas/getting-started) for development purposes. Create a new database named `VectorStore`, create a new collection in that database named `Embeddings`, and create an Atlas Vector Search index with the following JSON configuration:

```json
{
  "fields": [
    {
      "numDimensions": 1536,
      "path": "embedding",
      "similarity": "euclidean",
      "type": "vector"
    }
  ]
}
```

Add the Atlas connection information to the local `.env` file:

```bash
% cat >> demo/.env<< EOF
ATLAS_CONNECTION_STRING="mongodb+srv://<username>:<password>@vectorstore.ozmdcxv.mongodb.net/?retryWrites=false"
ATLAS_DB="VectorStore"
ATLAS_COLLECTION="Embeddings"
ATLAS_INDEX="vector_index"

EOF
```

## Verify complete `.env` file

If all of the prerequisites have been completed then the `.env` file should look something like this:

```bash
% cat demo/.env
REDPANDA_SERVERS="<bootstrap_server_url>"
REDPANDA_USER="<username>"
REDPANDA_PASS="<password>"
REDPANDA_TOPIC="documents"

OPENAI_API_KEY="<secret_key>"
OPENAI_EMBEDDING_MODEL="text-embedding-3-small"
OPENAI_MODEL="gpt-4o"

ATLAS_CONNECTION_STRING="mongodb+srv://<username>:<password>@vectorstore.ozmdcxv.mongodb.net/?retryWrites=false"
ATLAS_DB="VectorStore"
ATLAS_COLLECTION="Embeddings"
ATLAS_INDEX="vector_index"
```
