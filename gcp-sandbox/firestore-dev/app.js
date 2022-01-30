(async () => {
    try {
      const firebase = require('firebase')
      const config = {
        apiKey: 'xxx',
        // authDomain: '### FIREBASE AUTH DOMAIN ###',
        projectId: 'hoge',
      }
      firebase.initializeApp(config)
      const db = firebase.firestore()
  
      const userRef = db.collection('users').doc('xxx')
      const userDoc = await userRef.get()
      if (userDoc.exists) {
        console.log(userDoc.id)
        console.log(userDoc.data())
        console.log(userDoc.get('name'))
        console.log(userDoc.get('old'))
      } else {
        console.log('No such document!')
      }
      await db.app.delete()
    } catch (err) {
      console.log(`Error: ${JSON.stringify(err)}`)
    }
  })()