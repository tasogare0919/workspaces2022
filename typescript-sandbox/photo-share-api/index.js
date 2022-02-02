const { ApolloServer } = require('apollo-server')

const typeDefs = `
enum PhotoCategory {
    SELFIE
    PORTRAIT
    ACTION
    LANDSCAPE
    GRAPHIC
}

type User {
    githubLogin: ID!
    name: String
    avator: String
    postedPhotos: [Photo!]!
    inPhotos: [Photo!]!
}

type Photo {
    id: ID!
    url: String!
    name: String!
    description: String
    category: PhotoCategory!
    postedBy: User!
    taggedUsers: [User!]!
}

type Query {
    totalPhotos: Int!
    allPhotos: [Photo!]!
}

input PostPhotoInput {
    name: String!
    category: PhotoCategory=PORTRAIT
    description: String
}

type Mutation {
    postPhoto(name: String! description: String): Photo!
}
`

// 1.ユニークIDをインクリメントするための変数
var _id = 0
// 写真を格納するための配列
var photos = []
var tags = [
    { "photoID": "1", "userID": "gPlake"},
    { "photoID": "2", "userID": "sSchmidt" },
    { "photoID": "2", "userID": "mHattrup" },
    { "photoID": "2", "userID": "gPlake"}
]

const resolvers = {
    // 写真の枚数を確認できる
    Query: {
        totalPhotos: () => photos.length,
        allPhotos: () => photos
    },

    // postPhotoミューテーションと対応するリゾルバ
    Mutation : {
        postPhoto(parent, args) {
            // 2. 新しい写真を作成しidを作成する
            var newPhoto = {
                id: _id++,
                ...args
            }
            photos.push(newPhoto)
            // 3. 新しい写真を返す
            return newPhoto
        }
    }

    Photo: {
        url: parent => `http://yoursite.com/img/${parent.id}.jpg`,
        postedBy: parent => {
            return users.find(u => u.githubLogin === parent.githubUser)
        }
        taggedUsers: parent => tags
            .filter(tag => tag.photoID == parent.id)
            .map(tag => tag.userID)
            .map(userID => users.find(u => u.githubLogin === userID))
    }

    User: {
        postedPhotos: parent => {
            return photos.filter(p => p.githubUser === parent.githubLogin)
        }
        inPhotos: parent => tags
            .filter(tag => tag.userID == parent.id)
            .map(tag => tag.photoID)
            .map(photoID => photos.find(p => p.id === photoID))
    }
}

const server = new ApolloServer({
    typeDefs,
    resolvers
})

server.listen().then(({url}) => console.log(`GraphQL Service running on ${url}`))