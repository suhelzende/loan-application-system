$(function () {
  $.ajax({
    type: "GET",
    url: "http://localhost:8090/accounting/providers",
    contentType: "application/json",
    crossDomain: true,
    timeout: 1500,
    success: function (provider) {
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

    var usrname = $("#name").val();
    var usreamil = $("#email").val();
    var currentDate = new Date();

    loanApplication = {
      date: currentDate,
      borrower: {
        name: usrname,
        email: usreamil,
      },
    };

    $.ajax({
      type: "POST",
      url: "http://localhost:8090/application/start",
      contentType: "application/json",
      crossDomain: true,
      data: JSON.stringify(loanApplication),
      timeout: 1500,
      success: function (application) {
        loanApplication = application;
        $("#business").removeClass("hidden");
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
              alert("Failed to get data");
            },
          });
        }
      }
    }
  });

  $("#submit-application").submit(function (e) {
    e.preventDefault();

    $.ajax({
      type: "POST",
      url: "http://localhost:8090/application/submit",
      contentType: "application/json",
      crossDomain: true,
      data: JSON.stringify(loanDetails),
      timeout: 1500,
      success: function (application) {
        loanApplication = application;

        var status = "REJECTED";
        if ((application.Status = "ACCEPTED")) {
          status = "ACCEPTED";
        }

        outcome = ` 
            <div class="modal-content">
              <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal">&times;</button>
                <h4 class="modal-title">  ${status}</h4>
              </div>
              <div class="modal-body"> 
                <p>ApplicationID: ${loanApplication.ID}</p>
                <p>Business: ${loanApplication.BusinessDetails.Name}</p>
                <p>Loan Amount: ${loanApplication.LoanAmount}$</p>
                <p>Initiated Date: ${loanApplication.DateInitiated}.</p> 
                <p>Borrower : ${loanApplication.Borrower.name}</p>
                <p>Please keep this details for future reference</p>
              </div>
              <div class="modal-footer">
                <button type="button" id="closing" class="btn btn-default" data-dismiss="modal">Close</button>
              </div>
            </div> 
        `;
        document.getElementById("outcome").innerHTML = outcome;
        $("#result").modal();
      },
      dataType: "json",
      error: function () {
        alert("API failed");
      },
    });
  });

  $("#finalApproval").click(function () {
    $("#submitapplication").prop(
      "disabled",
      !$("#finalApproval").prop("checked")
    );
  });

  $("#result").on("hidden.bs.modal", function () {
    location.reload();
  });
});
