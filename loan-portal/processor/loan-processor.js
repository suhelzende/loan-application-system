$(function () {
  $.ajax({
    type: "GET",
    url: "http://localhost:8090/accounting/providers",
    contentType: "application/json",
    crossDomain: true,
    timeout: 1500,
    success: function (provider) {
      console.log(JSON.stringify(provider));
      providerOptions = provider
        .map((p) => `<option value="${p.ID}">${p.Name}</option>`)
        .join("");
      document.getElementById("providers").innerHTML =
        `<option value="0">Select Accounting Provider</option>` +
        providerOptions;
    },
    dataType: "json",
    error: function () {
      alert("error");
    },
  });

  var loanApplication;
  $("#start-application").submit(function (e) {
    e.preventDefault(); //prevent the default action

    console.log("submitted");
    var usrname = $("#name").val();
    var usreamil = $("#email").val();
    console.log(usrname);
    console.log(usreamil);
    var currentDate = new Date();
    console.log(currentDate);

    loanApplication = {
      date: currentDate,
      borrower: {
        name: usrname,
        email: usreamil,
      },
    };
    console.log(loanApplication);

    $.ajax({
      type: "POST",
      url: "http://localhost:8090/application/start",
      contentType: "application/json",
      crossDomain: true,
      data: JSON.stringify(loanApplication),
      timeout: 1500,
      success: function (application) {
        console.log(JSON.stringify(application));
        loanApplication = application;
      },
      dataType: "json",
      error: function () {
        alert("error");
      },
    });
  });

  $("#eYear").datepicker({
    format: "yyyy",
    viewMode: "years",
    minViewMode: "years",
    autoclose: true,
  });

  var loanDetails = {};
  $("#request").click(function (e) {
    e.preventDefault();
    if (loanApplication == undefined) {
      alert("Please start application first");
    } else {
      var businessName = $("#businessName").val();
      var businessID = $("#regId").val();
      var eYear = Number($("#eYear").val());
      var amount = Number($("#amount").val());
      console.log(businessID);
      console.log(businessName);
      console.log(eYear);
      console.log(amount);

      if (businessID == "" || businessName == "" || eYear == 0 || amount == 0) {
        alert("Please fill business details");
      } else {
        var accountProvider = $("#providers").val();
        if (accountProvider == 0) {
          alert("Please select accounting provider");
          return;
        } else {
          businessDetails = {
            registrationID: businessID,
            name: businessName,
            establishedYear: eYear,
          };

          accountingProvider = {
            id: accountProvider,
          };

          loanApplication.BusinessDetails = businessDetails;
          loanApplication.LoanAmount = amount;
          loanDetails.loanDetails = loanApplication;
          $.ajax({
            type: "POST",
            url: "http://localhost:8090/accounting/balencesheet/request",
            contentType: "application/json",
            crossDomain: true,
            data: JSON.stringify({
              businessDetails: businessDetails,
              accountingProvider: accountingProvider,
            }),
            timeout: 1500,
            success: function (sheet) {
              console.log(JSON.stringify(sheet));
              loanDetails.balanceSheet = sheet;
              balancesheet = sheet.Sheet.map(
                (s) =>
                  `<tr>
                  <td>${s.Year}</td>
                    <td>${s.Month}</td>
                    <td>${s.ProfitOrLoss}</td>
                    <td>${s.AssetsValue}</td>
                    </tr>`
              ).join("");
              document.getElementById("sheet").innerHTML = balancesheet;
              $("#balancesheettable").removeClass("hidden");
              $("#approval").removeClass("hidden");
            },
            dataType: "json",
            error: function () {
              alert("error");
            },
          });
        }
      }
    }
  });

  $("#submit-application").submit(function (e) {
    e.preventDefault(); //prevent the default action
    console.log(loanDetails);

    $.ajax({
      type: "POST",
      url: "http://localhost:8090/application/submit",
      contentType: "application/json",
      crossDomain: true,
      data: JSON.stringify(loanDetails),
      timeout: 1500,
      success: function (application) {
        console.log(JSON.stringify(application));
        loanApplication = application;

        outcome = `
        <div class="jumbotron jumbotron-fluid ">
          <div class="container">
            <h1 class="display-4">Rejected</h1>
            <p class="lead">Application has been rejected</p>
          </div>
        </div>
        `;
        if ((application.Status = "ACCEPTED")) {
          outcome = `
          <div class="jumbotron jumbotron-fluid" style="background:green;color:white">
            <div class="container" >
              <h1 class="display-4">Accpeted</h1>
              <p class="lead">Application has been accepeted.</p>
            </div>
          </div>
        `;
        }
        document.getElementById("outcome").innerHTML = outcome;
      },
      dataType: "json",
      error: function () {
        alert("error");
      },
    });
  });

  $("#finalApproval").click(function () {
    $("#submitapplication").prop(
      "disabled",
      !$("#finalApproval").prop("checked")
    );
  });
});
