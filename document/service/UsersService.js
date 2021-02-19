'use strict';


/**
 * Get all users' names
 * (description): Get all users' names
 *
 * returns List
 **/
exports.usersGET = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "firstname" : "Johnson",
  "lastname" : "Erick"
}, {
  "firstname" : "Johnson",
  "lastname" : "Erick"
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Delete a specified user
 * (description): Delete a specified user
 *
 * userId String user's id
 * returns Name
 **/
exports.usersUserIdDELETE = function(userId) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "firstname" : "Johnson",
  "lastname" : "Erick"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Get a specified user's name
 * (description): Get a specified user's name
 *
 * userId String user's id
 * returns Name
 **/
exports.usersUserIdGET = function(userId) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "firstname" : "Johnson",
  "lastname" : "Erick"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Create a new user
 * (description): Create a new user
 *
 * body Name (description): User's name
 * userId String user's id
 * no response value expected for this operation
 **/
exports.usersUserIdPOST = function(body,userId) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Update a specified user
 * (description): Update a specified user
 *
 * body Name (description): User's new name
 * userId String user's id
 * returns Name
 **/
exports.usersUserIdPUT = function(body,userId) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "firstname" : "Johnson",
  "lastname" : "Erick"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

