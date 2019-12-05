module.exports = { createTimestampedObject };

function createTimestampedObject(userContext, events, done) {
  const data = Date.now() ;
  // set the "data" variable for the virtual user to use in the subsequent action
  userContext.vars.data = data;
  return done();
}
