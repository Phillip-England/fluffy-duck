// client/utility/getById.ts
var getById = (elementId) => {
  const element = document.getElementById(elementId);
  if (element) {
    return element;
  }
  console.error(`failed to locate element by id: ${elementId}`);
  const div = document.createElement("div");
  div.innerHTML = `failed to located element by id: ${elementId}`;
  return document.createElement("div");
};

// client/events/eHookLoginForm.ts
var eHookLoginForm = () => {
  console.log("hooking event: eSubmitLoginForm");
  const form = getById("login-form");
  const error = getById("login-form-error");
  const hiddenTeamLink = getById("login-form-hidden-team-link");
  const hiddenAdminLink = getById("login-form-hidden-admin-link");
  const loader = getById("login-form-loader");
  const submit = getById("login-form-submit");
  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    submit.classList.add("hidden");
    loader.classList.remove("hidden");
    const data = new FormData(form);
    const username = data.get("username");
    const password = data.get("password");
    const res = await fetch("/api/user/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ username, password })
    });
    const json = await res.json();
    if (res.status == 200) {
      if (json.redirect == "/team") {
        hiddenTeamLink.click();
        return;
      }
      if (json.redirect == "/admin") {
        hiddenAdminLink.click();
        return;
      }
    }
    const errorMessage = json.message;
    error.innerHTML = errorMessage;
    error.classList.remove("hidden");
    submit.classList.remove("hidden");
    loader.classList.add("hidden");
  });
};

// client/utility/getByClass.ts
var getByClass = (elementClassName) => {
  const elements = document.getElementsByClassName(elementClassName);
  if (elements.length > 0) {
    return elements;
  }
  console.error(`failed to locate elements by class name: ${elementClassName}`);
  return `<p>failed to locate elements by class name: ${elementClassName}</p>`;
};

// client/events/eHookUpdateCemForm.ts
var eHookUpdateCemForm = async () => {
  console.log("hooking event: eUpdateCemForm");
  const cemUpdateSubmit = getById("update-cem-form-submit");
  const scoreRows = getByClass("cem-update-score-row");
  for (let i = 0;i < scoreRows.length; i++) {
    const row = scoreRows[i];
    const downButton = row.querySelector(".cem-score-update-down-arrow");
    const upButton = row.querySelector(".cem-score-update-up-arrow");
    const input = row.querySelector(".cem-score-update-input");
    downButton.addEventListener("click", () => {
      cemUpdateSubmit.scrollIntoView({ behavior: "smooth" });
      const currentScore = Number(input.value);
      if (currentScore - 1 > -1) {
        input.value = String(currentScore - 1);
      }
    });
    upButton.addEventListener("click", () => {
      cemUpdateSubmit.scrollIntoView({ behavior: "smooth" });
      const currentScore = Number(input.value);
      if (currentScore + 1 < 101) {
        input.value = String(currentScore + 1);
      }
    });
  }
};

// client/utility/toggleClassesOnElements.ts
function toggleClassesOnElements(elements, classes) {
  for (const element of elements) {
    for (const className of classes) {
      element.classList.toggle(className);
    }
  }
}

// client/events/eToggleNav.ts
var eToggleNav = () => {
  console.log("hooking event: eToggleNav");
  const bannerBars = getById("banner-bars");
  const bannerX = getById("banner-x");
  const nav = getById("nav-menu");
  const overlay = getById("overlay");
  bannerBars.addEventListener("click", () => {
    toggleClassesOnElements([bannerBars, bannerX, nav, overlay], ["hidden"]);
  });
  bannerX.addEventListener("click", () => {
    toggleClassesOnElements([bannerBars, bannerX, nav, overlay], ["hidden"]);
  });
  overlay.addEventListener("click", () => {
    toggleClassesOnElements([bannerBars, bannerX, nav, overlay], ["hidden"]);
  });
};

// client/index.ts
var clientRouter = new Map;
clientRouter.set("/", () => {
  console.log("/");
  eHookLoginForm();
});
clientRouter.set("/team", () => {
  console.log("/team");
  eToggleNav();
});
clientRouter.set("/admin", () => {
  console.log("/admin");
  eToggleNav();
});
clientRouter.set("/admin/cem/update", async () => {
  console.log("/admin/cem/update");
  eToggleNav();
  await eHookUpdateCemForm();
});
var clientHydrationFunc = clientRouter.get(window.location.pathname);
if (!clientHydrationFunc) {
  console.error(`no client hydration function set for route: ${window.location.pathname}`);
} else {
  clientHydrationFunc();
}
