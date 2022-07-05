import { render, screen } from '@testing-library/react';

import TodoItem from './TodoItem';

const handleToggle = jest.fn();

beforeEach(() => {
  handleToggle.mockReset();
});

describe('TodoItem', () => {
  it('renders active todo.', () => {
    render(
      <TodoItem todo={{ id: 2, title: 'Buy Unicorn', completed: false }} onToggle={handleToggle} />
    );

    const checkboxElement = screen.getByTestId('completed');
    expect(checkboxElement).not.toBeChecked();
  });

  it('renders completed todo.', () => {
    render(
      <TodoItem
        todo={{ id: 1, title: 'Taste JavaScript', completed: true }}
        onToggle={handleToggle}
      />
    );

    const checkboxElement = screen.getByTestId('completed');
    expect(checkboxElement).toBeChecked();
  });
});
