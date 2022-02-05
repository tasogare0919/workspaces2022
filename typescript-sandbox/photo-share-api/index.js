const express = require('express')
const { ApolloServer } = require('apollo-server-express')
const { MongoClient } = require('mongodb')
const { readFileSync } = require('fs')
const expressPlayground = require('graphql-playground-middleware-express').default
const resolvers = require(`./resolvers`)

require(`dotenv`).config()
const typeDefs = readFileSync(`./typeDefs.graphql`, `UTF-8`)


// ユニークIDをインクリメントするための変数
var _id = 0
// 写真を格納するための配列
var photos = []
var tags = [
    { "photoID": "1", "userID": "gPlake"},
    { "photoID": "2", "userID": "sSchmidt" },
    { "photoID": "2", "userID": "mHattrup" },
    { "photoID": "2", "userID": "gPlake"}
]

async function start() {
    const app = express()
    const MONGO_DB = process.env.DB_HOST
    let db

    try {
        const client = await MongoClient.connect(MONGO_DB, { useNewUrlParser: true })
        db = client.db()
    } catch(error) {
        console.log(`

            Mongo DB Host not found!
            please add DB_HOST environment variable to .env file

            existing...
        
        `)
    }
    process.exit(1)
}

const server = new ApolloServer({ 
    typeDefs,
    resolvers,
    context: async ({ req }) => {
        const githubToken = req.headers.authorization
        const currentUser = await db.connection('users').findOne({ githubToken })
        return { db, currentUser }
    }
})

// applyMiddlewareを呼び出しexpressにミドルウェアを追加
server.applyMiddleware( { app })


app.get('/playground', expressPlayground({ endpoint: '/graphql' }))
app.get('/', (req, res) => {
    let url = `https://github.com/login/oauth/authorize?client_id=${process.env.CLIENT_ID}&scope=user`
    res.end(`<a href="${url}>Sign In with Github</a>`)
})


app.listen({ port: 4000}, () =>
    console.log(`GraphQL server running at http://localhost:4000${server.graphqlPath}`)
)

start()
// const resolvers = {
//     // 写真の枚数を確認できる
//     Query: {
//         totalPhotos: () => photos.length,
//         allPhotos: () => photos
//     },

//     // postPhotoミューテーションと対応するリゾルバ
//     Mutation : {
//         postPhoto(parent, args) {
//             // 新しい写真を作成しidを作成する
//             var newPhoto = {
//                 id: _id++,
//                 ...args
//             }
//             photos.push(newPhoto)
//             // 新しい写真を返す
//             return newPhoto
//         }
//     }

//     Photo: {
//         url: parent => `http://yoursite.com/img/${parent.id}.jpg`,
//         postedBy: parent => {
//             return users.find(u => u.githubLogin === parent.githubUser)
//         }
//         taggedUsers: parent => tags
//             .filter(tag => tag.photoID == parent.id)
//             .map(tag => tag.userID)
//             .map(userID => users.find(u => u.githubLogin === userID))
//     }

//     User: {
//         postedPhotos: parent => {
//             return photos.filter(p => p.githubUser === parent.githubLogin)
//         }
//         inPhotos: parent => tags
//             .filter(tag => tag.userID == parent.id)
//             .map(tag => tag.photoID)
//             .map(photoID => photos.find(p => p.id === photoID))
//     }
// }

// const server = new ApolloServer({
//     typeDefs,
//     resolvers
// })
// server.listen().then(({url}) => console.log(`GraphQL Service running on ${url}`))