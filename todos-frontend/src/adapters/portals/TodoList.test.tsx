import { render, screen } from "@testing-library/react";

import TodoList from "./TodoList";

describe("TodoList", () => {
  it("renders all todos as active", () => {
    render(<TodoList activeCount={2} completedCount={0} />);

    const checkboxElement = screen.getByLabelText(/mark all as complete/i);
    expect(checkboxElement).not.toBeChecked();
  });

  it("renders all todos as completed", () => {
    render(<TodoList activeCount={0} completedCount={2} />);

    const checkboxElement = screen.getByLabelText(/mark all as complete/i);
    expect(checkboxElement).toBeChecked();
  });

  it("renders some todos as active or completed", () => {
    render(<TodoList activeCount={1} completedCount={1} />);

    const checkboxElement = screen.getByLabelText(/mark all as complete/i);
    expect(checkboxElement).toBePartiallyChecked();
  });
});
