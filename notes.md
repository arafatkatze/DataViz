worker.postMessage(params); // in the file index (webpage)
and the code is :      "src": editor.getSession().getDocument().getValue(),

Then we just insert one step to the process, i.e., first code is send to us, we get the code, then we just send back the data in specfic time interval, 
then the client will show the steps. It is easier for understanding.
better if the client can save the history and get the data step forward and backward. -> we can just add button for this. 
For server it just send back the data. The client can do the thing. No need to keep the long time connection?


Or we can use reflection, reflector, wrap a data structure, and then we inject the code that
is related to graph generation.
Then the data is collected. And we are good to go.

so we need a new interface, it collect information, print the result.
It should produce hooked version data. Then the data can be used the same way. 
but the data is collected.
and a structure that implements it.
