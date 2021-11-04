(async () => {
    try {
      const firebase = require('firebase')
      const config = {
        apiKey: 'AIzaSyAQdhn6qGQ6peEwC2a41w5wv9Qbo7Fob_Y',
        // authDomain: '### FIREBASE AUTH DOMAIN ###',
        projectId: 'kaggle-261706',
      }
      firebase.initializeApp(config)
      const db = firebase.firestore()
  
      const userRef = db.collection('users').doc('DYzRDWLmm1EiH628s8Pw')
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