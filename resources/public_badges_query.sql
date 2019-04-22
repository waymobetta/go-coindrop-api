select
	coindrop_tasks.badge_id
from
	coindrop_tasks
where
	coindrop_tasks.id = (
		select
			coindrop_user_tasks.task_id
		from
			coindrop_user_tasks
		where
			coindrop_user_tasks.completed = true
		and
			coindrop_user_tasks.user_id = (
				select
					coindrop_reddit.user_id
				from
					coindrop_reddit
				where
					coindrop_reddit.username = 'qa_adchain_registry'
			)
	)

