import { MyPromise } from "./my-promise.js";

const task = new MyPromise((resolve, reject) => {
    setTimeout(() => resolve(), 3000);
});

task.then(() => console.log("hhhhhhhhhhhhhh")).then(() => console.log(12345));
