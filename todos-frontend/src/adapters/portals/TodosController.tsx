import { useLocation } from "react-router-dom";
import { useState } from "react";

import {
  AddTodoCommand,
  ClearCompletedCommand,
  DestroyTodoCommand,
  SaveTodoCommand,
  SelectTodosQuery,
  SelectTodosQueryResult,
  ToggleAllCommand,
  ToggleTodoCommand,
} from "../../domain/messages";
import { Todo, TodoId } from "../../domain/data";

import { Filter } from "./utils";
import Footer from "./Footer";
import Header from "./Header";
import TodoItem from "./TodoItem";
import TodoList from "./TodoList";
import { useOnLoad } from "./hooks";

type TodosControllerProps = Readonly<{
  selectedTodos?: SelectTodosQueryResult;
  onAddTodo: (c: AddTodoCommand) => void;
  onClearCompleted: (c: ClearCompletedCommand) => void;
  onDestroyTodo: (c: DestroyTodoCommand) => void;
  onSaveTodo: (c: SaveTodoCommand) => void;
  onSelectTodos: (q: SelectTodosQuery) => void;
  onToggleAll: (c: ToggleAllCommand) => void;
  onToggleTodo: (c: ToggleTodoCommand) => void;
}>;

function TodosController({
  selectedTodos,
  onAddTodo,
  onClearCompleted,
  onDestroyTodo: onDestroy,
  onSaveTodo,
  onSelectTodos,
  onToggleAll,
  onToggleTodo,
}: TodosControllerProps) {
  const [editing, setEditing] = useState<TodoId | null>();

  function handleAddTodo(title: string) {
    onAddTodo({ title });
  }

  function handleToggleTodo(id: TodoId) {
    onToggleTodo({ id });
  }

  function handleToggleAll(checked: boolean) {
    onToggleAll({ checked });
  }

  function handleDestroy(id: TodoId) {
    onDestroy({ id });
  }

  function handleEdit(id: TodoId) {
    setEditing(id);
  }

  function handleSaveTodo(id: TodoId, newTitle: string) {
    onSaveTodo({ id, newTitle });
    setEditing(null);
  }

  function handleCancel() {
    setEditing(null);
  }

  function handleClearCompleted() {
    onClearCompleted({});
  }

  useOnLoad(() => onSelectTodos?.({}));

  const { activeCount, completedCount, filter, shownTodos } = useProjection(selectedTodos);

  return (
    <section className="relative container mx-auto mt-36 mb-10 bg-white shadow-2xl">
      <Header onAddTodo={handleAddTodo} />
      {selectedTodos?.todos.length ? (
        <>
          <TodoList
            activeCount={activeCount}
            completedCount={completedCount}
            onToggleAll={handleToggleAll}
          >
            {shownTodos.map((todo) => (
              <TodoItem
                key={todo.id}
                editing={editing === todo.id}
                todo={todo}
                onCancel={handleCancel}
                onDestroy={() => handleDestroy(todo.id)}
                onEdit={() => handleEdit(todo.id)}
                onSave={(title) => handleSaveTodo(todo.id, title)}
                onToggle={() => handleToggleTodo(todo.id)}
              />
            ))}
          </TodoList>
          <Footer
            activeCount={activeCount}
            completedCount={completedCount}
            filter={filter}
            onClearCompleted={handleClearCompleted}
          />
        </>
      ) : null}
    </section>
  );
}

export default TodosController;

type ProjectionResult = Readonly<{
  activeCount: number;
  completedCount: number;
  filter: Filter;
  shownTodos: readonly Todo[];
}>;

function useProjection(selectedTodos?: SelectTodosQueryResult): ProjectionResult {
  const { pathname } = useLocation();

  if (selectedTodos == null) {
    return {
      activeCount: 0,
      completedCount: 0,
      filter: Filter.All,
      shownTodos: [],
    };
  }

  let activeCount, completedCount, filter, shownTodos;
  switch (pathname) {
    case "/active":
      filter = Filter.Active;
      shownTodos = selectedTodos.todos.filter((t) => !t.completed);
      break;
    case "/completed":
      filter = Filter.Completed;
      shownTodos = selectedTodos.todos.filter((t) => t.completed);
      break;
    default:
      filter = Filter.All;
      shownTodos = selectedTodos.todos;
      break;
  }
  activeCount = selectedTodos.todos.filter((t) => !t.completed).length;
  completedCount = selectedTodos.todos.length - activeCount;

  return { activeCount, completedCount, filter, shownTodos };
}
