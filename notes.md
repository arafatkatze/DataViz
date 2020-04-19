worker.postMessage(params); // in the file index (webpage)
and the code is :      "src": editor.getSession().getDocument().getValue(),

Then we just insert one step to the process, i.e., first code is send to us, we get the code, then we just send back the data in specfic time interval, 
then the client will show the steps. It is easier for understanding.
better if the client can save the history and get the data step forward and backward. -> we can just add button for this. 
For server it just send back the data. The client can do the thing. No need to keep the long time connection?
