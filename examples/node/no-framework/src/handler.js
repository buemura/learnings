import { parse } from "node:url";
import { DEFAULT_HEADER } from "./utils/util.js";

const allRoutes = {
  "/heroes:get": async (request, response) => {
    throw new Error("TESTTTTTTT");
    response.write("GET");
    response.end();
  },
  // 404 routes
  default: (request, response) => {
    response.writeHead(404, DEFAULT_HEADER);
    response.write("Not found");
    response.end();
  },
};

function handler(request, response) {
  const { url, method } = request;
  const { pathname } = parse(url, true);
  const key = `${pathname}:${method.toLowerCase()}`;
  const chosen = allRoutes[key] || allRoutes.default;

  return Promise.resolve(chosen(request, response)).catch(
    errorHandler(response)
  );
}

function errorHandler(response) {
  return (error) => {
    console.log("Something bad has happened**", error.stack);
    response.writeHead(500, DEFAULT_HEADER);
    response.write(
      JSON.stringify({
        error: "Internal server error",
      })
    );

    return response.end();
  };
}

export default handler;
