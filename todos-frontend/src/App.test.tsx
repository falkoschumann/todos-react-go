import { render, screen } from "@testing-library/react";
import { Router } from "react-router-dom";
import { createMemoryHistory } from "history";

import App from "./App";

describe("App", () => {
  test("renders successfully.", () => {
    const history = createMemoryHistory();
    history.push("/");

    render(
      <Router location={history.location} navigator={history}>
        <App />
      </Router>
    );

    const linkElement = screen.getByText(/todos/i);
    expect(linkElement).toBeInTheDocument();
  });
});
