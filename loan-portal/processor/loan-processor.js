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

    $.ajax({
      type: "POST",
      url: "http://localhost:8090/application/start",
      contentType: "application/json",
      crossDomain: true,
      data: JSON.stringify(body),
      timeout: 1500,
      success: function (msg) {
        console.log(JSON.stringify(msg));
        $("#businesstab").addClass("active").attr("aria-expanded", "true");
        $("#hometab").removeClass("active").attr("aria-expanded", "false");
      },
      dataType: "json",
      error: function () {
        alert("error");
      },
    });
  });
});
