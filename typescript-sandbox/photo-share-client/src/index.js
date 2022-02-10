import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import ApolloClient , { gql } from 'apollo-boost'
import { ApolloProvider } from 'react-apollo';

const client = new ApolloClient({ uri: 'http://localhost:4000/graphql' })
const query = gql`
{
  totalUsers
  totalPhotos
}
`
client.query({query}).then(({ data }) => console.log('data', data)).catch(console.error)

render (
  <ApolloProvider client={client}>
    <App />
  </ApolloProvider>,
  document.getElementById('root')
)