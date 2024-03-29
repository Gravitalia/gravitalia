defmodule Notification do
    @moduledoc """
    Notification allows access to Gravitalia's in-app notifications using several technologies
    """

  use Application
  require Logger

  def start(_type, _args) do
    children = [
      Plug.Cowboy.child_spec(
        scheme: :http,
        plug: Notification.Router,
        options: [
          port: Application.fetch_env!(:notification, :port)
        ],
        protocol_options: [idle_timeout: :infinity]
      ),
      Registry.child_spec(
        keys: :duplicate,
        name: Registry.Notification
      ),
      {PubSub, []}
    ]

    Logger.info("Server started at http://localhost:#{Application.fetch_env!(:notification, :port)}")

    opts = [strategy: :one_for_one, name: Notification.Application]
    Supervisor.start_link(children, opts)
  end
end
