# AutoDevGPT

Automatically create software development agents

Rather than having 1 generic agent trying to solve a very vauge problem we should at least setup the agent for some success by providing it a framework in which to complete a
task. By utilizing the traditional software development workflow & utilizing multiple agents with different prompts we are able to produce higher quality output.

* We are implementing an adversairal agent whos task is to prune & provide feedback on tasks to better steer the overall project.
* AI dont have eyes so its hard for them to see if any artistic output is any good, so we are also adding in a human feedback loop whereby the user can "nudge" the worker in
  the direction they want while it is going through its tasks.


* [ ] Use adversairial agent to validate if task is complete
* [ ] Implement plugin action system to facilitate extending the base Agents with more functionality
* [ ] Implement software development workflow agents
* [ ] Implement human feedback loop on tasks where the AI Agent needs human clarification (images etc)