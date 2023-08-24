$(function () {
  $("#start-application").submit(function (e) {
    e.preventDefault(); //prevent the default action

    console.log("submitted");
    var usrname = $("#name").val();
    var usreamil = $("#email").val();
    console.log(usrname);
    console.log(usreamil);
    var currentDate = new Date();
    console.log(currentDate);

    var body = {
      date: currentDate,
      borrower: {
        name: usrname,
        email: usreamil,
      },
    };
    console.log(body);
    fetch("http://localhost:8090/application/start", {
      method: "POST", // *GET, POST, PUT, DELETE, etc.
      mode: "no-cors", // no-cors, *cors, same-origin
      cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
      credentials: "same-origin", // include, *same-origin, omit
      headers: {
        "Content-Type": "application/json",
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      redirect: "follow", // manual, *follow, error
      referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
      body: JSON.stringify(body), // body data type must match "Content-Type" header
    }).then((response) => {
      console.log(response);
    });
  });
});
